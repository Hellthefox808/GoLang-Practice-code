// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"strings"
	"testing"
)

func TestDecodeStrictJSON(t *testing.T) {
	valid := []byte(`{"id":"c1","title":"Go Testing","price_usd":10,"created_at":"2026-11-17T00:00:00Z"}`)
	c, err := DecodeStrictJSON(valid)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Title != "Go Testing" {
		t.Errorf("unexpected title: %s", c.Title)
	}

	// Unknown field check
	invalid := []byte(`{"id":"c2","title":"Go Testing","unknown_field":true}`)
	_, err = DecodeStrictJSON(invalid)
	if err == nil || !strings.Contains(err.Error(), "unknown field") {
		t.Errorf("expected unknown field error, got %v", err)
	}
}
