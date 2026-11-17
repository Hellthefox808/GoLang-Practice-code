// Author: Ravi Ranjan Singh
// License: MIT

package main

import "testing"

func TestSumArray(t *testing.T) {
	arr := [4]int{5, 10, 15, 20}
	expected := 50
	if got := SumArray(arr); got != expected {
		t.Errorf("SumArray() = %d, expected %d", got, expected)
	}
}

func TestInvertMatrix2x2(t *testing.T) {
	matrix := [2][2]int{
		{1, 2},
		{3, 4},
	}
	expected := [2][2]int{
		{1, 3},
		{2, 4},
	}
	got := InvertMatrix2x2(matrix)
	if got != expected {
		t.Errorf("InvertMatrix2x2() = %v, expected %v", got, expected)
	}
}
