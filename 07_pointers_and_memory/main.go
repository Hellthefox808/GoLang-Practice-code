// Package main demonstrates pointer semantics in Go, memory address referencing (&),
// value dereferencing (*), pass-by-value fundamentals, heap vs stack escape analysis,
// and nil pointer defenses.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"errors"
	"fmt"
)

var (
	ErrNilPointer = errors.New("cannot perform operation on nil pointer")
)

type ServerConfig struct {
	Host string
	Port int
	TLS  bool
}

// MutateByValue receives a copy of the config; mutations do NOT affect the caller.
func MutateByValue(cfg ServerConfig, newPort int) {
	cfg.Port = newPort
}

// MutateByPointer receives a copy of the pointer address; mutations modify the underlying heap/stack data.
func MutateByPointer(cfg *ServerConfig, newPort int) error {
	if cfg == nil {
		return ErrNilPointer
	}
	cfg.Port = newPort
	return nil
}

// SafeIncrement safely increments the integer value pointed to.
func SafeIncrement(val *int, delta int) error {
	if val == nil {
		return ErrNilPointer
	}
	*val += delta
	return nil
}

func main() {
	fmt.Println("=== 07: Pointers, Memory Model, and Mutation Semantics ===")

	// 1. Basic Pointers & Addresses
	count := 42
	var ptr *int = &count

	fmt.Printf("Variable 'count'  : Value=%d, Memory Address=%p\n", count, &count)
	fmt.Printf("Pointer 'ptr'     : Holds Address=%p, Points To Value=%d\n\n", ptr, *ptr)

	// 2. Dereferencing & Mutation
	*ptr = 100
	fmt.Printf("After '*ptr = 100': 'count' is now %d\n\n", count)

	// 3. Pass-by-value vs Pointer semantics on Structs
	cfg := ServerConfig{
		Host: "localhost",
		Port: 8080,
		TLS:  false,
	}

	fmt.Println("Initial Config:", cfg)
	MutateByValue(cfg, 9090)
	fmt.Println("After MutateByValue (Should be unchanged):", cfg)

	err := MutateByPointer(&cfg, 9090)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("After MutateByPointer (Should be 9090)   :", cfg)
	}

	// 4. Nil Pointer Defense
	var uninitializedPtr *ServerConfig
	err = MutateByPointer(uninitializedPtr, 7070)
	if errors.Is(err, ErrNilPointer) {
		fmt.Println("\nDefensive Programming: Successfully intercepted nil pointer mutation attempt.")
	}

	// 5. Safe Increment
	counter := 10
	_ = SafeIncrement(&counter, 5)
	fmt.Printf("Counter incremented safely to: %d\n", counter)
}
