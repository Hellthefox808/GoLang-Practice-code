// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"errors"
	"testing"
	"unsafe"
)

func TestNewUser(t *testing.T) {
	u, err := NewUser("u1", "Ravi Ranjan Singh", "ravi@example.com", 25)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.ID != "u1" || u.FullName != "Ravi Ranjan Singh" || !u.IsActive {
		t.Errorf("unexpected user values: %+v", u)
	}

	_, err = NewUser("u2", "Test", "bademail", 20)
	if !errors.Is(err, ErrInvalidEmail) {
		t.Errorf("expected ErrInvalidEmail, got %v", err)
	}

	_, err = NewUser("u3", "Test", "valid@mail.com", 16)
	if !errors.Is(err, ErrUnderageUser) {
		t.Errorf("expected ErrUnderageUser, got %v", err)
	}
}

func TestStructAlignmentOptimization(t *testing.T) {
	unoptSize := unsafe.Sizeof(UnoptimizedStruct{})
	optSize := unsafe.Sizeof(OptimizedStruct{})

	if optSize >= unoptSize {
		t.Errorf("expected optimized struct size (%d) to be smaller than unoptimized (%d)", optSize, unoptSize)
	}
}
