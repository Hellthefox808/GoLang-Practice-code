// Package main examines fixed-size arrays in Go, their value semantics,
// type identity governed by length [N]T, multi-dimensional array layouts, and iteration.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import "fmt"

// SumArray computes the sum of a fixed 4-element integer array.
// Notice that [4]int cannot accept a [5]int; length is strictly part of the type signature.
func SumArray(arr [4]int) int {
	total := 0
	for _, val := range arr {
		total += val
	}
	return total
}

// InvertMatrix2x2 transposes a fixed 2x2 integer matrix.
func InvertMatrix2x2(matrix [2][2]int) [2][2]int {
	return [2][2]int{
		{matrix[0][0], matrix[1][0]},
		{matrix[0][1], matrix[1][1]},
	}
}

func main() {
	fmt.Println("=== 08: Fixed-Size Arrays & Value Semantics ===")

	// 1. Array declaration with explicit size
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[2] = "Orange"
	fruitList[3] = "Peach"

	fmt.Printf("Fruit List      : %v (Length: %d)\n", fruitList, len(fruitList))

	// 2. Composite literal with ellipsis operator (...) for automatic sizing
	numbers := [...]int{10, 20, 30, 40}
	fmt.Printf("Numbers Array   : %v (Computed Type: %T)\n", numbers, numbers)
	fmt.Printf("Sum of Numbers  : %d\n\n", SumArray(numbers))

	// 3. Array Value Copy Semantics (Arrays are NOT references; they are values!)
	original := [3]string{"Go", "Rust", "C++"}
	copied := original // Fully copies all elements!
	copied[0] = "Python"

	fmt.Println("--- Array Value Copy Verification ---")
	fmt.Printf("Original Array  : %v\n", original)
	fmt.Printf("Copied Array    : %v (Original remains completely unmodified)\n\n", copied)

	// 4. Multidimensional Arrays (2D Matrix)
	matrix := [2][2]int{
		{1, 2},
		{3, 4},
	}
	transposed := InvertMatrix2x2(matrix)
	fmt.Println("Original 2x2 Matrix   :", matrix)
	fmt.Println("Transposed 2x2 Matrix :", transposed)
}
