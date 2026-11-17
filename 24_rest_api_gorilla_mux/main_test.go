// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRESTAPIEndpoints(t *testing.T) {
	repo := NewCourseRepository()
	ctrl := NewController(repo)
	router := BuildServerRouter(ctrl)

	// 1. Test GET /api/courses
	t.Run("GET /api/courses", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/courses", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200 OK, got %d", rr.Code)
		}

		var resp StandardAPIResponse
		_ = json.Unmarshal(rr.Body.Bytes(), &resp)
		if !resp.Success {
			t.Errorf("expected success true in response envelope")
		}
	})

	// 2. Test GET /api/courses/101
	t.Run("GET /api/courses/101", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/courses/101", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Fatalf("expected 200 OK, got %d", rr.Code)
		}
	})

	// 3. Test GET /api/courses/999 (Not Found)
	t.Run("GET /api/courses/999 Not Found", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/courses/999", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("expected 404 NotFound, got %d", rr.Code)
		}
	})

	// 4. Test POST /api/courses (Create)
	t.Run("POST /api/courses Create", func(t *testing.T) {
		payload := []byte(`{"title":"Kubernetes Production Architecture","price_usd":199.0}`)
		req := httptest.NewRequest(http.MethodPost, "/api/courses", bytes.NewBuffer(payload))
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusCreated {
			t.Fatalf("expected 201 Created, got %d: %s", rr.Code, rr.Body.String())
		}
	})

	// 5. Test DELETE /api/courses/101
	t.Run("DELETE /api/courses/101", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/api/courses/101", nil)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusNoContent {
			t.Errorf("expected 204 NoContent, got %d", rr.Code)
		}
	})
}
