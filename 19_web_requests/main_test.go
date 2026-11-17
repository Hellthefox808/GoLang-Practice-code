// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFetchData(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("mock-data-payload"))
	}))
	defer ts.Close()

	client := &http.Client{Timeout: time.Second}
	data, code, err := FetchData(context.Background(), client, ts.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if code != http.StatusOK || !strings.Contains(data, "mock-data-payload") {
		t.Errorf("unexpected response: code=%d, body=%s", code, data)
	}
}
