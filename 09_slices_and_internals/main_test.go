// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"errors"
	"slices"
	"testing"
)

func TestRemoveAtIndexOrderPreserving(t *testing.T) {
	orig := []string{"A", "B", "C", "D"}
	res, err := RemoveAtIndexOrderPreserving(orig, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []string{"A", "C", "D"}
	if !slices.Equal(res, expected) {
		t.Errorf("got %v, expected %v", res, expected)
	}

	_, err = RemoveAtIndexOrderPreserving(orig, 10)
	if !errors.Is(err, ErrIndexOutOfBounds) {
		t.Errorf("expected ErrIndexOutOfBounds, got %v", err)
	}
}

func TestRemoveAtIndexFast(t *testing.T) {
	orig := []int{10, 20, 30, 40}
	res, err := RemoveAtIndexFast(orig, 1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Index 1 (20) replaced by last element (40) -> [10, 40, 30]
	expected := []int{10, 40, 30}
	if !slices.Equal(res, expected) {
		t.Errorf("got %v, expected %v", res, expected)
	}
}
