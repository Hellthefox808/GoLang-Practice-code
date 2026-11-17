// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"errors"
	"io"
	"strings"
	"testing"
)

func TestReadCleanLine(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    string
		expectedErr error
	}{
		{
			name:        "Unix line feed",
			input:       "Go Developer\n",
			expected:    "Go Developer",
			expectedErr: nil,
		},
		{
			name:        "Windows CRLF",
			input:       "Ravi Ranjan Singh\r\n",
			expected:    "Ravi Ranjan Singh",
			expectedErr: nil,
		},
		{
			name:        "Empty input with LF",
			input:       "   \n",
			expected:    "",
			expectedErr: ErrEmptyInput,
		},
		{
			name:        "EOF on empty string",
			input:       "",
			expected:    "",
			expectedErr: io.EOF,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadCleanLine(strings.NewReader(tt.input))
			if tt.expectedErr != nil {
				if !errors.Is(err, tt.expectedErr) {
					t.Fatalf("expected error %v, got %v", tt.expectedErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("got %q, expected %q", got, tt.expected)
			}
		})
	}
}

func TestCollectTokens(t *testing.T) {
	input := "Go Microservices Docker"
	tokens := CollectTokens(strings.NewReader(input))
	if len(tokens) != 3 {
		t.Fatalf("expected 3 tokens, got %d", len(tokens))
	}
	if tokens[0] != "Go" || tokens[1] != "Microservices" || tokens[2] != "Docker" {
		t.Errorf("unexpected tokens: %v", tokens)
	}
}
