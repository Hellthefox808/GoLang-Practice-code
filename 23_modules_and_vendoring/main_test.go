// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBuildRouter(t *testing.T) {
	router := BuildRouter()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}

	// Verify security headers middleware
	if rr.Header().Get("X-Content-Type-Options") != "nosniff" {
		t.Error("expected nosniff header from SecurityHeadersMiddleware")
	}
}
