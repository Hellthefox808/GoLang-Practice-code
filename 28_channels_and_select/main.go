// Package main details Go Channels, CSP communication, directional channel constraints,
// the select multiplexer, non-blocking operations, and production Worker Pool architectures.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID       int
	WorkData string
}

type Result struct {
	JobID  int
	Output string
	Err    error
}

// Worker processes jobs from jobsChan and emits completed results to resultsChan.
// Notice strict directionality: jobsChan is receive-only, resultsChan is send-only.
func Worker(id int, jobsChan <-chan Job, resultsChan chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobsChan {
		// Simulate processing latency
		time.Sleep(10 * time.Millisecond)
		resultsChan <- Result{
			JobID:  job.ID,
			Output: fmt.Sprintf("Worker #%d successfully processed: %s", id, job.WorkData),
			Err:    nil,
		}
	}
}

// RunWorkerPool orchestrates a fixed worker pool consuming a stream of tasks.
func RunWorkerPool(numWorkers int, jobs []Job) []Result {
	jobsChan := make(chan Job, len(jobs))
	resultsChan := make(chan Result, len(jobs))
	var wg sync.WaitGroup

	// 1. Spawn workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go Worker(w, jobsChan, resultsChan, &wg)
	}

	// 2. Dispatch jobs into queue
	for _, j := range jobs {
		jobsChan <- j
	}
	// Close jobsChan to notify workers that no more tasks will be enqueued
	close(jobsChan)

	// 3. Wait for all workers to finish on a separate coordinator goroutine
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// 4. Collect results
	var collected []Result
	for res := range resultsChan {
		collected = append(collected, res)
	}

	return collected
}

func main() {
	fmt.Println("=== 28: Channels, Multiplexing Select, and Worker Pools ===")

	// 1. Buffered vs Unbuffered Channels
	unbuffered := make(chan string)
	go func() {
		unbuffered <- "Synchronous Handshake Message"
	}()
	msg := <-unbuffered
	fmt.Printf("Unbuffered Sync Received: %s\n\n", msg)

	// 2. Select Statement Multiplexing with Timeout Protection
	fastChan := make(chan string, 1)
	slowChan := make(chan string, 1)

	go func() {
		time.Sleep(20 * time.Millisecond)
		fastChan <- "Fast Event Dispatched"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		slowChan <- "Slow Event Dispatched"
	}()

	fmt.Println("--- Select Multiplexing Demonstration ---")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	select {
	case fastVal := <-fastChan:
		fmt.Printf("Intercepted First: %s\n", fastVal)
	case slowVal := <-slowChan:
		fmt.Printf("Intercepted Slow: %s\n", slowVal)
	case <-ctx.Done():
		fmt.Println("Timeout triggered! Operation exceeded 100ms.")
	}

	// 3. Worker Pool Architecture
	fmt.Println("\n--- Executing 3-Worker Pool on 6 Jobs ---")
	var jobs []Job
	for i := 1; i <= 6; i++ {
		jobs = append(jobs, Job{
			ID:       i,
			WorkData: fmt.Sprintf("High-Concurrency Data Payload #%d", i),
		})
	}

	results := RunWorkerPool(3, jobs)
	for _, r := range results {
		fmt.Printf("  [Result Job %d]: %s\n", r.JobID, r.Output)
	}
}
