// Package main explores explicit type conversions, string-to-numeric parsing,
// numeric-to-string formatting via strconv, numeric truncation risks, and safe bounds checking.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	ErrIntegerOverflow = errors.New("integer exceeds target bit-width boundary")
)

// ParseRating extracts and validates a floating-point score between min and max bounds.
func ParseRating(input string, minVal, maxVal float64) (float64, error) {
	trimmed := strings.TrimSpace(input)
	val, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid numeric representation %q: %w", input, err)
	}

	if val < minVal || val > maxVal {
		return 0, fmt.Errorf("value %.2f out of acceptable range [%.1f - %.1f]", val, minVal, maxVal)
	}

	return val, nil
}

// SafeInt64ToInt32 checks numeric boundaries before downcasting int64 to int32 to prevent overflow.
func SafeInt64ToInt32(val int64) (int32, error) {
	if val > math.MaxInt32 || val < math.MinInt32 {
		return 0, fmt.Errorf("%w: value %d does not fit in int32", ErrIntegerOverflow, val)
	}
	return int32(val), nil
}

func main() {
	fmt.Println("=== 04: Type Conversions, Parsing, and Safe Casts ===")

	// 1. String to Numeric parsing
	ratingStr := " 4.85 \n"
	rating, err := ParseRating(ratingStr, 1.0, 5.0)
	if err != nil {
		fmt.Printf("Rating parsing failed: %v\n", err)
	} else {
		fmt.Printf("Parsed Rating: %.2f (Bonus applied: %.2f)\n", rating, rating+0.15)
	}

	// 2. String to Integer
	intStr := "4096"
	port, err := strconv.Atoi(intStr)
	if err != nil {
		fmt.Printf("Port conversion error: %v\n", err)
	} else {
		fmt.Printf("Server Port: %d (Hex format: 0x%X)\n", port, port)
	}

	// 3. Boolean parsing and formatting
	boolStr := "true"
	parsedBool, _ := strconv.ParseBool(boolStr)
	formattedBool := strconv.FormatBool(parsedBool)
	fmt.Printf("Bool parsed: %t, formatted back: %s\n\n", parsedBool, formattedBool)

	// 4. Safe Downcasting vs Unsafe Overflow
	fmt.Println("--- Boundary-Safe Downcasting ---")
	var normalInt64 int64 = 65535
	asInt32, err := SafeInt64ToInt32(normalInt64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Safe Cast Result: %d (int32)\n", asInt32)
	}

	var overflowInt64 int64 = math.MaxInt64
	_, err = SafeInt64ToInt32(overflowInt64)
	if err != nil {
		fmt.Printf("Expected Overflow Caught: %v\n", err)
	}

	// 5. Byte Slice and String memory conversions
	message := "Go Concurrency Mastery by Ravi Ranjan Singh"
	byteSlice := []byte(message)
	fmt.Printf("\nOriginal String: %s\nByte Slice Len : %d bytes\nReconstructed  : %s\n",
		message, len(byteSlice), string(byteSlice))
}
