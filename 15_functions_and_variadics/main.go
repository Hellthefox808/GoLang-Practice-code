// Package main explores function mechanics in Go: first-class functions,
// closures with lexical state, variadic arguments, multiple return signatures,
// and functional abstractions (Map / Filter).
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("cannot divide by zero")

// VariadicSum accepts any number of integers, including zero.
func VariadicSum(numbers ...int) (total int, count int) {
	for _, n := range numbers {
		total += n
	}
	return total, len(numbers)
}

// SafeDivide demonstrates multiple return values with explicit error handling.
func SafeDivide(numerator, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, ErrDivideByZero
	}
	return numerator / denominator, nil
}

// SequenceGenerator returns a closure function encapsulating private mutable state.
func SequenceGenerator(initial int, step int) func() int {
	current := initial
	return func() int {
		val := current
		current += step
		return val
	}
}

// GenericMap applies a transformation function to each element in a slice.
func GenericMap[T, U any](input []T, transform func(T) U) []U {
	result := make([]U, len(input))
	for i, v := range input {
		result[i] = transform(v)
	}
	return result
}

func main() {
	fmt.Println("=== 15: Functions, Variadic Signatures, and Closures ===")

	// 1. Variadic function with direct arguments and slice unpacking
	total, count := VariadicSum(10, 20, 30, 40)
	fmt.Printf("Variadic direct  : Sum=%d (from %d numbers)\n", total, count)

	scores := []int{100, 200, 300}
	totalScore, scoreCount := VariadicSum(scores...) // Unpacking slice via '...'
	fmt.Printf("Variadic unpacked: Sum=%d (from %d numbers)\n\n", totalScore, scoreCount)

	// 2. Safe division with multiple returns
	quotient, err := SafeDivide(100, 4)
	if err == nil {
		fmt.Printf("100 / 4 = %.2f\n", quotient)
	}

	_, err = SafeDivide(100, 0)
	if errors.Is(err, ErrDivideByZero) {
		fmt.Println("Division by zero cleanly intercepted without panic.")
	}

	// 3. Stateful Closures
	nextID := SequenceGenerator(1000, 5)
	fmt.Println("\nClosure ID generation:")
	fmt.Println("  ID 1:", nextID())
	fmt.Println("  ID 2:", nextID())
	fmt.Println("  ID 3:", nextID())

	// 4. Higher-Order Functions with Generics
	rawStrings := []string{"golang", "microservices", "docker"}
	uppercased := GenericMap(rawStrings, func(s string) string {
		return fmt.Sprintf("DEV-%s", s)
	})
	fmt.Printf("\nTransformed via GenericMap: %v\n", uppercased)
}
