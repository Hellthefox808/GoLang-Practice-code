// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"math/big"
	"testing"
)

func TestGenerateSecureToken(t *testing.T) {
	token, err := GenerateSecureToken(16)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// 16 bytes encoded in hex produces 32 characters
	if len(token) != 32 {
		t.Errorf("expected token length 32, got %d (%s)", len(token), token)
	}

	_, err = GenerateSecureToken(0)
	if err == nil {
		t.Error("expected error for byteSize <= 0, got nil")
	}
}

func TestGenerateCryptoInt(t *testing.T) {
	for i := 0; i < 20; i++ {
		val, err := GenerateCryptoInt(10)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if val < 0 || val >= 10 {
			t.Errorf("value %d out of bounds [0, 10)", val)
		}
	}
}

func TestFinancialPrecisionAdd(t *testing.T) {
	sum, err := FinancialPrecisionAdd("0.1", "0.2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected, _, _ := big.ParseFloat("0.3", 10, 128, big.ToNearestEven)
	if sum.Cmp(expected) != 0 {
		t.Errorf("expected 0.3, got %s", sum.Text('f', 10))
	}
}
