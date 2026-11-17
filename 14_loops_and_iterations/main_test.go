// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"slices"
	"testing"
)

func TestFilterEvenSquares(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	expected := []int{4, 16, 36}
	got := FilterEvenSquares(input)

	if !slices.Equal(got, expected) {
		t.Errorf("got %v, expected %v", got, expected)
	}
}

func TestFindMatrixTarget(t *testing.T) {
	grid := [][]int{
		{1, 2},
		{3, 4},
	}
	r, c, found := FindMatrixTarget(grid, 4)
	if !found || r != 1 || c != 1 {
		t.Errorf("expected found at (1,1), got found=%t at (%d,%d)", found, r, c)
	}

	_, _, found = FindMatrixTarget(grid, 99)
	if found {
		t.Error("expected false for missing target")
	}
}
