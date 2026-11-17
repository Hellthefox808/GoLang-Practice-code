// Package main explores modern Go module architecture, dependency hygiene,
// standard library Go 1.22+ enhanced routing, middleware pipelines,
// and graceful HTTP server termination.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// LoggingMiddleware records request latency, HTTP method, and path.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s %s - %v", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
	})
}

// SecurityHeadersMiddleware injects enterprise defense headers into all responses.
func SecurityHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		next.ServeHTTP(w, r)
	})
}

// BuildRouter configures routes using Go 1.22+ enhanced HTTP multiplexing.
func BuildRouter() http.Handler {
	mux := http.NewServeMux()

	// Go 1.22+ supports HTTP method prefixes and path parameters directly in stdlib!
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"service": "Go Modules & Middleware Showcase",
			"author":  "Ravi Ranjan Singh",
			"status":  "healthy",
		})
	})

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"UP"}`))
	})

	// Wrap router through middleware stack
	handler := SecurityHeadersMiddleware(mux)
	handler = LoggingMiddleware(handler)

	return handler
}

func main() {
	fmt.Println("=== 23: Go Modules, Standard Router (Go 1.22+), and Graceful Shutdown ===")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      BuildRouter(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to capture operating system interruption signals (SIGINT, SIGTERM)
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	// Launch server asynchronously in a background goroutine
	go func() {
		fmt.Printf("HTTP Server listening on http://localhost:%s\n", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server startup failure: %v\n", err)
		}
	}()

	// In an interactive demo or background service, listen for shutdown trigger
	// For testing, trigger shutdown after a short demonstration delay if non-interactive
	go func() {
		time.Sleep(100 * time.Millisecond)
		stopChan <- os.Interrupt
	}()

	<-stopChan
	fmt.Println("\nShutdown signal intercepted. Initiating graceful drain...")

	// 5-second deadline context for in-flight requests to finish
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server drained and stopped gracefully. Zero connection drops.")
}
