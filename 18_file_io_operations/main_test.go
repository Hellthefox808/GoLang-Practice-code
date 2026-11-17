// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAtomicWriteAndReadLines(t *testing.T) {
	tmpDir := t.TempDir()
	target := filepath.Join(tmpDir, "test.txt")

	data := []byte("alpha\nbeta\ngamma\n")
	err := AtomicWriteFile(target, data, 0644)
	if err != nil {
		t.Fatalf("unexpected write error: %v", err)
	}

	lines, err := ReadLinesChunked(target)
	if err != nil {
		t.Fatalf("unexpected read error: %v", err)
	}

	if len(lines) != 3 || lines[0] != "alpha" || lines[1] != "beta" || lines[2] != "gamma" {
		t.Errorf("unexpected lines read: %v", lines)
	}
}
