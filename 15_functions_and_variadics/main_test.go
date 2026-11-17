// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"errors"
	"slices"
	"testing"
)

func TestVariadicSum(t *testing.T) {
	tot, count := VariadicSum(1, 2, 3, 4, 5)
	if tot != 15 || count != 5 {
		t.Errorf("got (%d, %d), expected (15, 5)", tot, count)
	}

	totEmpty, countEmpty := VariadicSum()
	if totEmpty != 0 || countEmpty != 0 {
		t.Errorf("expected 0, 0 for empty call")
	}
}

func TestSafeDivide(t *testing.T) {
	q, err := SafeDivide(10, 2)
	if err != nil || q != 5.0 {
		t.Errorf("expected 5.0, got %f (err: %v)", q, err)
	}

	_, err = SafeDivide(10, 0)
	if !errors.Is(err, ErrDivideByZero) {
		t.Errorf("expected ErrDivideByZero, got %v", err)
	}
}

func TestSequenceGenerator(t *testing.T) {
	gen := SequenceGenerator(10, 2)
	if gen() != 10 || gen() != 12 || gen() != 14 {
		t.Error("sequence generator generated unexpected progression")
	}
}

func TestGenericMap(t *testing.T) {
	in := []int{1, 2, 3}
	out := GenericMap(in, func(n int) int { return n * 10 })
	expected := []int{10, 20, 30}
	if !slices.Equal(out, expected) {
		t.Errorf("got %v, expected %v", out, expected)
	}
}
