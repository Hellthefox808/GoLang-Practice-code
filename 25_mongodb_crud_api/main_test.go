
/ Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Hellthefox808/GoLang-Practice-code/25_mongodb_crud_api/controller"
	"github.com/Hellthefox808/GoLang-Practice-code/25_mongodb_crud_api/router"
)

func TestMovieRouterDegraded(t *testing.T) {
	// Controller with nil collection gracefully responds without crashing
	ctrl := controller.NewMovieController(nil)
	r := router.SetupRouter(ctrl)

	// GET /api/movies
	req := httptest.NewRequest(http.MethodGet, "/api/movies", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}

	// POST /api/movies with empty title
	reqPost := httptest.NewRequest(http.MethodPost, "/api/movies", bytes.NewBuffer([]byte(`{"title":""}`)))
	rrPost := httptest.NewRecorder()
	r.ServeHTTP(rrPost, reqPost)

	if rrPost.Code != http.StatusUnprocessableEntity {
		t.Errorf("expected 422 UnprocessableEntity, got %d", rrPost.Code)
	}
}
