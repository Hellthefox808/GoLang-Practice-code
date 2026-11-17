// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"strings"
	"testing"
)

func TestParseAndInspectURL(t *testing.T) {
	u, q, err := ParseAndInspectURL("https://example.com/search?q=golang&limit=10")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if u.Scheme != "https" || u.Host != "example.com" || u.Path != "/search" {
		t.Errorf("unexpected url components: %v", u)
	}

	if q.Get("q") != "golang" || q.Get("limit") != "10" {
		t.Errorf("unexpected query values: %v", q)
	}
}

func TestBuildAPIEndpoint(t *testing.T) {
	res, err := BuildAPIEndpoint("https://api.test.com/base/", "subpath", map[string]string{
		"token": "xyz123",
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if !strings.HasPrefix(res, "https://api.test.com/base/subpath") || !strings.Contains(res, "token=xyz123") {
		t.Errorf("unexpected endpoint string: %s", res)
	}
}
