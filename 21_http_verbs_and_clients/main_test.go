// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
)

func TestDispatchVerbs(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("GET_OK"))
		case http.MethodPost:
			if strings.Contains(r.Header.Get("Content-Type"), "json") {
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte("POST_JSON_OK"))
			} else {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("POST_FORM_OK"))
			}
		}
	}))
	defer server.Close()

	client := &http.Client{Timeout: 2 * time.Second}
	ctx := context.Background()

	// Test GET
	getRes, err := DispatchGet(ctx, client, server.URL)
	if err != nil || getRes != "GET_OK" {
		t.Errorf("GET failed: %s, %v", getRes, err)
	}

	// Test POST JSON
	postJSONRes, err := DispatchPostJSON(ctx, client, server.URL, map[string]string{"foo": "bar"})
	if err != nil || postJSONRes != "POST_JSON_OK" {
		t.Errorf("POST JSON failed: %s, %v", postJSONRes, err)
	}

	// Test POST Form
	form := url.Values{"key": {"val"}}
	postFormRes, err := DispatchPostForm(ctx, client, server.URL, form)
	if err != nil || postFormRes != "POST_FORM_OK" {
		t.Errorf("POST Form failed: %s, %v", postFormRes, err)
	}
}
