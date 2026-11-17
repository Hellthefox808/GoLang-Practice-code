// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBoundedHealthChecker(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer ts.Close()

	urls := []string{ts.URL, ts.URL, ts.URL}
	res := BoundedHealthChecker(urls, 2)

	if len(res) != 3 {
		t.Fatalf("expected 3 results, got %d", len(res))
	}

	for _, r := range res {
		if r.StatusCode != http.StatusOK || r.Err != nil {
			t.Errorf("unexpected health check failure: %+v", r)
		}
	}
}
