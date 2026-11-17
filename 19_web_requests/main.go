// Package main demonstrates production-hardened HTTP client architecture in Go:
// custom timeouts, context deadline propagation, transport pool reuse,
// and complete response body draining to prevent socket connection leaks.
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
	"time"
)

// HTTPClientConfig defines enterprise timeout budgets for outbound network calls.
type HTTPClientConfig struct {
	Timeout time.Duration
}

// FetchData retrieves string content from an endpoint using a context-bounded client.
func FetchData(ctx context.Context, client *http.Client, targetURL string) (string, int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, targetURL, nil)
	if err != nil {
		return "", 0, fmt.Errorf("failed to construct request: %w", err)
	}

	req.Header.Set("User-Agent", "GoLang-Practice-Suite/2.0 (Ravi-Ranjan-Singh)")
	req.Header.Set("Accept", "text/plain, application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", 0, fmt.Errorf("request dispatch failed: %w", err)
	}

	// Always ensure the response body is drained and closed to return the TCP socket
	// to the internal net/http Transport connection keep-alive pool.
	defer func() {
		_, _ = io.Copy(io.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", resp.StatusCode, fmt.Errorf("failed reading response payload: %w", err)
	}

	return string(bodyBytes), resp.StatusCode, nil
}

func main() {
	fmt.Println("=== 19: Production HTTP Clients & Connection Pool Reuse ===")

	// 1. Spinning up a local deterministic in-memory test server
	// Guarantees this practice code runs offline without reliance on flaky third-party URLs!
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"success","author":"Ravi Ranjan Singh","module":"19_web_requests"}`))
	}))
	defer server.Close()

	// 2. Production HTTP Client with strict timeout boundaries
	customClient := &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        100,
			MaxIdleConnsPerHost: 10,
			IdleConnTimeout:     90 * time.Second,
		},
	}

	// 3. Dispatch request with a 2-second context deadline
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Printf("Dispatching GET request to in-memory endpoint: %s\n", server.URL)
	payload, statusCode, err := FetchData(ctx, customClient, server.URL)
	if err != nil {
		fmt.Printf("HTTP Request failed: %v\n", err)
		return
	}

	fmt.Printf("Response Status Code: %d OK\n", statusCode)
	fmt.Printf("Response Payload    : %s\n\n", payload)

	// 4. Timeout handling demonstration
	slowServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond) // Simulated latency
		w.WriteHeader(http.StatusOK)
	}))
	defer slowServer.Close()

	shortCtx, shortCancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer shortCancel()

	fmt.Println("Dispatching request with an intentionally short deadline (50ms):")
	_, _, err = FetchData(shortCtx, customClient, slowServer.URL)
	if err != nil {
		fmt.Printf("Verified Deadline Interception: %v\n", err)
	}
}
