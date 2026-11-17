// Package main explores slice internals, backing arrays, SliceHeader anatomy (ptr, len, cap),
// reallocation dynamics during append, safe sub-slicing, and modern slices package algorithms.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"errors"
	"fmt"
	"slices"
)

var (
	ErrIndexOutOfBounds = errors.New("slice index out of bounds")
)

// RemoveAtIndexOrderPreserving removes an element at index, shifting subsequent elements left (O(N)).
func RemoveAtIndexOrderPreserving[T any](s []T, index int) ([]T, error) {
	if index < 0 || index >= len(s) {
		return nil, ErrIndexOutOfBounds
	}
	return append(s[:index], s[index+1:]...), nil
}

// RemoveAtIndexFast removes an element at index by swapping with the last element (O(1), order altered).
func RemoveAtIndexFast[T any](s []T, index int) ([]T, error) {
	if index < 0 || index >= len(s) {
		return nil, ErrIndexOutOfBounds
	}
	lastIdx := len(s) - 1
	s[index] = s[lastIdx]
	return s[:lastIdx], nil
}

func main() {
	fmt.Println("=== 09: Slices, Capacity Growth, and Memory Internals ===")

	// 1. Slice creation via make with predefined length and capacity
	// Pre-allocating capacity avoids recurrent memory allocations and GC overhead
	metrics := make([]int, 0, 5)
	fmt.Printf("Initial Allocation: len=%d, cap=%d\n", len(metrics), cap(metrics))

	// Appending within capacity
	for i := 1; i <= 5; i++ {
		metrics = append(metrics, i*10)
		fmt.Printf("Append %d -> len=%d, cap=%d\n", i*10, len(metrics), cap(metrics))
	}

	// Appending beyond capacity triggers backing array reallocation and copy
	metrics = append(metrics, 60)
	fmt.Printf("Append 60 (Exceeded Cap) -> len=%d, cap=%d (New Backing Array!)\n\n", len(metrics), cap(metrics))

	// 2. Sub-slicing and shared backing array demonstration
	data := []string{"Go", "Rust", "TypeScript", "Zig", "C++"}
	subSlice := data[1:4] // "Rust", "TypeScript", "Zig"
	fmt.Printf("Original: %v\nSubSlice: %v\n", data, subSlice)

	// Modifying subSlice alters original data because they share the backing array!
	subSlice[0] = "RUST_MUTATED"
	fmt.Printf("After mutating subSlice[0]: Original is now %v\n\n", data)

	// 3. Independent clone via slices.Clone (Go 1.21+)
	cloned := slices.Clone(data)
	cloned[0] = "INDEPENDENT"
	fmt.Printf("Cloned Slice: %v\nOriginal unaffected: %v\n\n", cloned, data)

	// 4. Removing elements
	frameworks := []string{"Gin", "Fiber", "Echo", "Chi", "Gorilla"}
	fmt.Println("Initial Frameworks:", frameworks)

	orderedRemoved, _ := RemoveAtIndexOrderPreserving(frameworks, 2) // Removes "Echo"
	fmt.Println("Order Preserved (Removed 'Echo'):", orderedRemoved)

	fastRemoved, _ := RemoveAtIndexFast(orderedRemoved, 0) // Fast remove index 0
	fmt.Println("Fast O(1) Removal (Unordered)  :", fastRemoved)

	// 5. Modern slices algorithms (Go 1.21+)
	numbers := []int{45, 12, 85, 32, 89, 39, 69, 44, 42, 1, 45, 23}
	slices.Sort(numbers)
	fmt.Printf("\nSorted Numbers: %v\n", numbers)
	fmt.Printf("Contains 42?  : %t\n", slices.Contains(numbers, 42))
}
