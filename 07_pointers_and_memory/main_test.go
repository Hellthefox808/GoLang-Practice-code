// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"errors"
	"testing"
)

func TestMutateByPointer(t *testing.T) {
	cfg := &ServerConfig{Host: "127.0.0.1", Port: 80, TLS: false}

	err := MutateByPointer(cfg, 443)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.Port != 443 {
		t.Errorf("expected port 443, got %d", cfg.Port)
	}

	err = MutateByPointer(nil, 8080)
	if !errors.Is(err, ErrNilPointer) {
		t.Errorf("expected ErrNilPointer, got %v", err)
	}
}

func TestSafeIncrement(t *testing.T) {
	val := 50
	err := SafeIncrement(&val, 25)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 75 {
		t.Errorf("expected val 75, got %d", val)
	}

	err = SafeIncrement(nil, 10)
	if !errors.Is(err, ErrNilPointer) {
		t.Errorf("expected ErrNilPointer, got %v", err)
	}
}
