// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestGreet(t *testing.T) {
	tests := []struct {
		name        string
		inputName   string
		expectedSub string
	}{
		{
			name:        "Custom name greeting",
			inputName:   "Ravi Ranjan Singh",
			expectedSub: "Hello, Ravi Ranjan Singh!",
		},
		{
			name:        "Default name fallback",
			inputName:   "",
			expectedSub: "Hello, Go Developer!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			n, err := Greet(&buf, tt.inputName)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if n <= 0 {
				t.Errorf("expected positive byte count written, got %d", n)
			}
			output := buf.String()
			if !strings.Contains(output, tt.expectedSub) {
				t.Errorf("expected output to contain %q, got %q", tt.expectedSub, output)
			}
		})
	}
}
