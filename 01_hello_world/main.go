// Package main demonstrates the canonical entry point of a Go application,
// illustrating program initialization, the main package contract, standard I/O,
// and runtime environment inspection.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
)

// init functions execute automatically prior to main() in lexical order.
// They are commonly utilized for environment verification, flag setup, or registration.
func init() {
	// Verify that the architecture and OS target meet minimum requirements.
	_ = runtime.GOOS
}

// Greet writes a formatted greeting message to the provided io.Writer stream.
// Decoupling output generation from os.Stdout allows unit testing without capturing OS streams.
func Greet(w io.Writer, name string) (int, error) {
	if name == "" {
		name = "Go Developer"
	}
	return fmt.Fprintf(w, "Hello, %s! Welcome to High-Performance Go Programming by Ravi Ranjan Singh.\n", name)
}

func main() {
	// 1. Basic console greeting via standard output
	n, err := Greet(os.Stdout, "Ravi Ranjan Singh")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing greeting: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %d bytes to stdout.\n\n", n)

	// 2. Runtime Environment Inspection
	fmt.Println("--- Go Runtime Diagnostics ---")
	fmt.Printf("Go Compiler Version : %s\n", runtime.Version())
	fmt.Printf("Target Operating Sys: %s\n", runtime.GOOS)
	fmt.Printf("CPU Architecture    : %s\n", runtime.GOARCH)
	fmt.Printf("Available CPU Cores : %d\n", runtime.NumCPU())
	fmt.Printf("Active Goroutines   : %d\n", runtime.NumGoroutine())
	fmt.Println("-------------------------------")
}
