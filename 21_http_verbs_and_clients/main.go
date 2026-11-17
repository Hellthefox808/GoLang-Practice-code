// Package main demonstrates dispatching diverse HTTP verbs (GET, POST JSON, POST Form, PUT, DELETE)
// with custom request headers, body streams, status checks, and self-contained mock server verification.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"time"
)

// DispatchGet executes an HTTP GET request with custom headers.
func DispatchGet(ctx context.Context, client *http.Client, endpoint string) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// DispatchPostJSON serializes a data payload and posts it as application/json.
func DispatchPostJSON(ctx context.Context, client *http.Client, endpoint string, payload any) (string, error) {
	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("json serialization failed: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// DispatchPostForm posts form values as application/x-www-form-urlencoded.
func DispatchPostForm(ctx context.Context, client *http.Client, endpoint string, formData url.Values) (string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	fmt.Println("=== 21: HTTP Verbs (GET, POST JSON, Form-Encoded, Custom Headers) ===")

	// Spin up an in-process mock server handling GET, POST JSON, and POST Form routes
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/api/get":
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(`{"message":"GET Success from Mock Server","author":"Ravi Ranjan Singh"}`))
		case "/api/post":
			var received map[string]any
			_ = json.NewDecoder(r.Body).Decode(&received)
			resp, _ := json.Marshal(map[string]any{
				"status":   "received_json",
				"payload":  received,
				"auth_header": r.Header.Get("Authorization"),
			})
			w.WriteHeader(http.StatusCreated)
			_, _ = w.Write(resp)
		case "/api/postform":
			_ = r.ParseForm()
			resp, _ := json.Marshal(map[string]any{
				"status": "received_form",
				"fields": r.PostForm,
			})
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(resp)
		default:
			http.NotFound(w, r)
		}
	}))
	defer mockServer.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	ctx := context.Background()

	// 1. Dispatch GET
	fmt.Println("--- 1. HTTP GET ---")
	getResp, err := DispatchGet(ctx, client, mockServer.URL+"/api/get")
	if err != nil {
		fmt.Println("GET Error:", err)
	} else {
		fmt.Println("GET Response:", getResp)
	}

	// 2. Dispatch POST JSON
	fmt.Println("\n--- 2. HTTP POST JSON ---")
	coursePayload := map[string]any{
		"course":    "Go Distributed Systems",
		"instructor": "Ravi Ranjan Singh",
		"duration":  "40 Hours",
	}
	postJSONResp, err := DispatchPostJSON(ctx, client, mockServer.URL+"/api/post", coursePayload)
	if err != nil {
		fmt.Println("POST JSON Error:", err)
	} else {
		fmt.Println("POST JSON Response:", postJSONResp)
	}

	// 3. Dispatch POST Form
	fmt.Println("\n--- 3. HTTP POST Form (x-www-form-urlencoded) ---")
	form := url.Values{}
	form.Set("username", "raviranjansingh")
	form.Set("role", "Lead Architect")
	form.Set("status", "Active")

	postFormResp, err := DispatchPostForm(ctx, client, mockServer.URL+"/api/postform", form)
	if err != nil {
		fmt.Println("POST Form Error:", err)
	} else {
		fmt.Println("POST Form Response:", postFormResp)
	}
}
