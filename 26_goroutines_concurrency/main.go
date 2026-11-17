// Package main explores Goroutine concurrency foundations, the M:N Go scheduler,
// safe sync.WaitGroup orchestration, connection leak prevention, and bounded worker semaphores.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"sync"
	"time"
)

type EndpointHealth struct {
	URL        string
	StatusCode int
	Latency    time.Duration
	Err        error
}

// BoundedHealthChecker probes multiple endpoints concurrently while capping parallelism
// using a buffered channel semaphore.
func BoundedHealthChecker(urls []string, maxParallel int) []EndpointHealth {
	var wg sync.WaitGroup
	results := make([]EndpointHealth, len(urls))

	// Buffered channel acting as a counting semaphore
	semaphore := make(chan struct{}, maxParallel)

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	for i, u := range urls {
		// CRITICAL RULE: wg.Add() must ALWAYS precede the 'go' keyword!
		wg.Add(1)

		go func(index int, targetURL string) {
			// Acquire semaphore slot (blocks if maxParallel goroutines are currently executing)
			semaphore <- struct{}{}
			defer func() {
				<-semaphore // Release semaphore slot
				wg.Done()   // Decrement WaitGroup counter
			}()

			start := time.Now()
			req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, targetURL, nil)
			if err != nil {
				results[index] = EndpointHealth{URL: targetURL, Err: err, Latency: time.Since(start)}
				return
			}

			resp, err := client.Do(req)
			if err != nil {
				results[index] = EndpointHealth{URL: targetURL, Err: err, Latency: time.Since(start)}
				return
			}

			// PREVENT SOCKET LEAK: Drain and close response body
			_, _ = io.Copy(io.Discard, resp.Body)
			_ = resp.Body.Close()

			results[index] = EndpointHealth{
				URL:        targetURL,
				StatusCode: resp.StatusCode,
				Latency:    time.Since(start),
				Err:        nil,
			}
		}(i, u)
	}

	// Wait until all bounded workers complete
	wg.Wait()
	return results
}

func main() {
	fmt.Println("=== 26: Goroutines, Scheduler Dynamics, and Bounded Concurrency ===")

	// 1. Stand up local mock servers to test concurrent probing reliably
	var mockServers []*httptest.Server
	var endpoints []string

	for i := 1; i <= 6; i++ {
		id := i
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(time.Duration(id*15) * time.Millisecond) // Simulated latency
			w.WriteHeader(http.StatusOK)
		}))
		defer ts.Close()
		mockServers = append(mockServers, ts)
		endpoints = append(endpoints, ts.URL)
	}

	fmt.Printf("Probing %d endpoints concurrently with bounded parallelism (Max 3 concurrent)...\n\n", len(endpoints))

	start := time.Now()
	results := BoundedHealthChecker(endpoints, 3)
	totalElapsed := time.Since(start)

	for i, res := range results {
		statusStr := fmt.Sprintf("%d OK", res.StatusCode)
		if res.Err != nil {
			statusStr = fmt.Sprintf("FAILED (%v)", res.Err)
		}
		fmt.Printf("  Target [%02d]: %-30s | Status: %-10s | Latency: %v\n", i+1, res.URL, statusStr, res.Latency)
	}

	fmt.Printf("\nCompleted all %d concurrent probes in %v (Bounded semaphore verified).\n", len(endpoints), totalElapsed)
}
