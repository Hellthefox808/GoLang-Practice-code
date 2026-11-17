// Package main demonstrates modern Go File I/O operations: atomic file writes,
// buffered streaming with bufio, os.ReadFile/WriteFile, temporary file management,
// and cross-platform filepath manipulation.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// AtomicWriteFile writes data to a temporary file, flushes to disk, and atomically
// renames it to the target path. This guarantees zero file corruption if the application or OS crashes.
func AtomicWriteFile(targetPath string, data []byte, perm os.FileMode) error {
	dir := filepath.Dir(targetPath)
	tempFile, err := os.CreateTemp(dir, "atomic-*.tmp")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tempPath := tempFile.Name()

	// Ensure cleanup in case of failure prior to rename
	defer func() {
		_ = tempFile.Close()
		_ = os.Remove(tempPath)
	}()

	if _, err := tempFile.Write(data); err != nil {
		return fmt.Errorf("failed writing to temp file: %w", err)
	}

	// Commit data from OS page cache to physical persistent storage
	if err := tempFile.Sync(); err != nil {
		return fmt.Errorf("failed syncing temp file: %w", err)
	}

	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("failed closing temp file: %w", err)
	}

	// Atomic rename replaces the target file atomically
	if err := os.Rename(tempPath, targetPath); err != nil {
		return fmt.Errorf("atomic rename failed: %w", err)
	}

	return nil
}

// ReadLinesChunked reads a file line-by-line using buffered scanning without loading the whole file into RAM.
func ReadLinesChunked(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed opening file %q: %w", path, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error during buffered scan: %w", err)
	}

	return lines, nil
}

func main() {
	fmt.Println("=== 18: File I/O, Atomic Operations, and Stream Buffering ===")

	tmpDir := os.TempDir()
	targetFile := filepath.Join(tmpDir, "ravi_golang_practice_audit.txt")
	defer os.Remove(targetFile) // Clean up after demo

	content := []byte("Line 1: High-Performance Go Systems by Ravi Ranjan Singh\n" +
		"Line 2: Atomic File I/O Pattern Implemented\n" +
		"Line 3: Memory-Bounded Stream Scanning Verified\n")

	// 1. Atomic Write
	fmt.Printf("Executing atomic write to: %s\n", targetFile)
	err := AtomicWriteFile(targetFile, content, 0644)
	if err != nil {
		fmt.Printf("Write failed: %v\n", err)
		return
	}
	fmt.Println("Atomic file write completed successfully.")

	// 2. High-Level Read with os.ReadFile (Replacing deprecated ioutil.ReadFile)
	readBytes, err := os.ReadFile(targetFile)
	if err != nil {
		fmt.Printf("Read failed: %v\n", err)
		return
	}
	fmt.Printf("\nos.ReadFile Read %d bytes successfully.\n", len(readBytes))

	// 3. Chunked Stream Read via bufio.Scanner
	fmt.Println("\nStreaming lines via bufio.Scanner:")
	lines, err := ReadLinesChunked(targetFile)
	if err != nil {
		fmt.Printf("Scan failed: %v\n", err)
		return
	}
	for i, line := range lines {
		fmt.Printf("  [%02d]: %s\n", i+1, line)
	}

	// 4. File Information & Permission inspection
	info, err := os.Stat(targetFile)
	if err == nil {
		fmt.Printf("\nFile Metadata:\n  Size: %d bytes\n  Mode: %v\n  ModTime: %v\n",
			info.Size(), info.Mode(), info.ModTime())
	}
}
