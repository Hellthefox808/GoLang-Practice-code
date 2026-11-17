// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"errors"
	"math"
	"testing"
)

func TestParseRating(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		minVal      float64
		maxVal      float64
		expectedVal float64
		expectErr   bool
	}{
		{"Valid rating", "4.5", 1.0, 5.0, 4.5, false},
		{"Valid with whitespace", "  3.8 \n", 1.0, 5.0, 3.8, false},
		{"Below minimum", "0.5", 1.0, 5.0, 0, true},
		{"Above maximum", "5.5", 1.0, 5.0, 0, true},
		{"Invalid non-numeric", "excellent", 1.0, 5.0, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val, err := ParseRating(tt.input, tt.minVal, tt.maxVal)
			if (err != nil) != tt.expectErr {
				t.Fatalf("expected error=%v, got err=%v", tt.expectErr, err)
			}
			if !tt.expectErr && math.Abs(val-tt.expectedVal) > 1e-6 {
				t.Errorf("got %f, expected %f", val, tt.expectedVal)
			}
		})
	}
}

func TestSafeInt64ToInt32(t *testing.T) {
	t.Run("Valid cast within bounds", func(t *testing.T) {
		got, err := SafeInt64ToInt32(1000)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if got != 1000 {
			t.Errorf("expected 1000, got %d", got)
		}
	})

	t.Run("Overflow beyond MaxInt32", func(t *testing.T) {
		_, err := SafeInt64ToInt32(int64(math.MaxInt32) + 1)
		if !errors.Is(err, ErrIntegerOverflow) {
			t.Fatalf("expected ErrIntegerOverflow, got %v", err)
		}
	})
}
