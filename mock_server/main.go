// Package main provides a lightweight, pure Go HTTP mock server for testing
// web request verbs, JSON payloads, and form submissions locally on port 8000.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("=== Local HTTP Mock Server for Go Web Requests ===")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Welcome to Go Mock Server authored by Ravi Ranjan Singh"))
	})

	mux.HandleFunc("GET /get", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]any{
			"message": "Hello from Go Mock Server",
			"author":  "Ravi Ranjan Singh",
			"status":  "healthy",
		})
	})

	mux.HandleFunc("POST /post", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Failed to read body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(bodyBytes)
	})

	mux.HandleFunc("POST /postform", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		_ = json.NewEncoder(w).Encode(r.PostForm)
	})

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Mock server active at http://localhost:%s\n", port)
	log.Fatal(server.ListenAndServe())
}
