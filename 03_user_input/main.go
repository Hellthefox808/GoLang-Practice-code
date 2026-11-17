// Package main explores idiomatic user and stream input handling in Go using
// bufio.Reader and bufio.Scanner, dealing with multi-platform newline delimiters
// (\r\n vs \n), input trimming, and stream termination.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	ErrEmptyInput = errors.New("input cannot be empty or solely whitespace")
)

// ReadCleanLine reads a single line from any io.Reader, strips both Windows CRLF (\r\n)
// and Unix LF (\n), and validates that the content is non-empty.
func ReadCleanLine(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	raw, err := reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", fmt.Errorf("failed to read from input stream: %w", err)
	}

	// Sanitize trailing CR and LF safely
	cleaned := strings.TrimRight(raw, "\r\n")
	cleaned = strings.TrimSpace(cleaned)

	if cleaned == "" {
		if errors.Is(err, io.EOF) {
			return "", io.EOF
		}
		return "", ErrEmptyInput
	}

	return cleaned, nil
}

// CollectTokens utilizes bufio.Scanner with bufio.ScanWords to extract discrete tokens.
func CollectTokens(r io.Reader) []string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	var tokens []string
	for scanner.Scan() {
		tokens = append(tokens, scanner.Text())
	}
	return tokens
}

func main() {
	fmt.Println("=== 03: Robust User & Stream Input Handling ===")
	fmt.Print("Please enter your developer name: ")

	name, err := ReadCleanLine(os.Stdin)
	if err != nil {
		if errors.Is(err, ErrEmptyInput) {
			fmt.Println("Warning: Defaulting developer name to 'Ravi Ranjan Singh'")
			name = "Ravi Ranjan Singh"
		} else {
			fmt.Fprintf(os.Stderr, "Fatal input error: %v\n", err)
			return
		}
	}

	fmt.Printf("Registered Developer: %q (Length: %d runes)\n\n", name, len([]rune(name)))

	// Example demonstration with simulated token scanning
	sampleStream := strings.NewReader("Golang Concurrency Goroutines Microservices Kubernetes")
	fmt.Println("Extracting tokens from stream:")
	tokens := CollectTokens(sampleStream)
	for i, token := range tokens {
		fmt.Printf("  Token [%02d]: %s\n", i+1, token)
	}
}
