// Author: Ravi Ranjan Singh
// License: MIT

package main

import "testing"

func TestLookupLanguage(t *testing.T) {
	registry := map[string]string{
		"GO": "Golang",
		"RS": "Rust",
	}

	val, ok := LookupLanguage(registry, "GO")
	if !ok || val != "Golang" {
		t.Errorf("expected Golang, true; got %s, %t", val, ok)
	}

	_, ok = LookupLanguage(registry, "UNKNOWN")
	if ok {
		t.Errorf("expected false for UNKNOWN key, got true")
	}
}
