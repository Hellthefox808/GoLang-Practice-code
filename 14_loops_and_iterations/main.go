// Package main demonstrates iterative control structures in Go: 3-component loops,
// condition loops, range iterations over slices/maps/runes, labeled breaks,
// and Go 1.22+ per-iteration variable scoping.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import "fmt"

// FilterEvenSquares returns squares of even integers using standard looping.
func FilterEvenSquares(nums []int) []int {
	var result []int
	for _, n := range nums {
		if n%2 == 0 {
			result = append(result, n*n)
		}
	}
	return result
}

// FindMatrixTarget searches a 2D grid using a labeled break to escape nested iteration immediately upon match.
func FindMatrixTarget(grid [][]int, target int) (row, col int, found bool) {
SearchLoop:
	for r, rowSlice := range grid {
		for c, val := range rowSlice {
			if val == target {
				row = r
				col = c
				found = true
				break SearchLoop // Breaks out of the outer loop immediately!
			}
		}
	}
	return row, col, found
}

func main() {
	fmt.Println("=== 14: Loops, Range Iterators, and Labeled Jumps ===")

	// 1. Classic 3-component loop
	sum := 0
	for i := 1; i <= 5; i++ {
		sum += i
	}
	fmt.Printf("Sum 1..5: %d\n", sum)

	// 2. While-style condition loop
	countdown := 3
	for countdown > 0 {
		fmt.Printf("Countdown: %d...\n", countdown)
		countdown--
	}
	fmt.Println("Liftoff!")

	// 3. Range iteration over UTF-8 String (Runes vs Byte Offsets)
	// Notice that multi-byte Unicode characters advance byte indices by more than 1!
	unicodeText := "Go-Go"
	fmt.Printf("\nIterating over Unicode string: %q\n", unicodeText)
	for byteOffset, charRune := range unicodeText {
		fmt.Printf("  Byte offset %2d: Character %c (Unicode codepoint: %U)\n", byteOffset, charRune, charRune)
	}

	// 4. Labeled Break out of nested matrix search
	grid := [][]int{
		{10, 20, 30},
		{40, 50, 60},
		{70, 80, 90},
	}
	r, c, ok := FindMatrixTarget(grid, 50)
	fmt.Printf("\nSearching 50 in Matrix: Found=%t at [row=%d, col=%d]\n", ok, r, c)

	// 5. Go 1.22+ Per-Iteration Scoping demonstration
	// In Go 1.22+, each iteration creates a new variable instance, fixing closure capture bugs!
	fmt.Println("\nGo 1.22+ Closure Iteration Scoping:")
	var funcs []func()
	values := []string{"Alpha", "Beta", "Gamma"}

	for _, v := range values {
		funcs = append(funcs, func() {
			fmt.Printf("  Captured closure value: %s\n", v)
		})
	}

	for _, f := range funcs {
		f() // Each prints its respective value without manual re-binding!
	}
}
