// Package main demonstrates modern Go mathematical operations, secure cryptographic
// random generation, PRNG vs CSPRNG distinctions, and arbitrary-precision arithmetic with math/big.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	cryptoRand "crypto/rand"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	mathRand "math/rand"
	"time"
)

// GenerateSecureToken generates a cryptographically secure hex token of the specified byte size.
func GenerateSecureToken(byteSize int) (string, error) {
	if byteSize <= 0 {
		return "", fmt.Errorf("byteSize must be greater than zero, got %d", byteSize)
	}
	bytes := make([]byte, byteSize)
	_, err := cryptoRand.Read(bytes)
	if err != nil {
		return "", fmt.Errorf("failed to generate secure random bytes: %w", err)
	}
	return hex.EncodeToString(bytes), nil
}

// GenerateCryptoInt generates a uniform random integer in [0, max) using CSPRNG.
func GenerateCryptoInt(max int64) (int64, error) {
	if max <= 0 {
		return 0, fmt.Errorf("max must be positive, got %d", max)
	}
	n, err := cryptoRand.Int(cryptoRand.Reader, big.NewInt(max))
	if err != nil {
		return 0, fmt.Errorf("failed to generate crypto int: %w", err)
	}
	return n.Int64(), nil
}

// FinancialPrecisionAdd adds two precise monetary amounts using math/big to prevent float inaccuracy.
func FinancialPrecisionAdd(amountA, amountB string) (*big.Float, error) {
	a, _, err := big.ParseFloat(amountA, 10, 128, big.ToNearestEven)
	if err != nil {
		return nil, fmt.Errorf("invalid amount A: %w", err)
	}
	b, _, err := big.ParseFloat(amountB, 10, 128, big.ToNearestEven)
	if err != nil {
		return nil, fmt.Errorf("invalid amount B: %w", err)
	}
	result := new(big.Float).SetPrec(128).Add(a, b)
	return result, nil
}

func main() {
	fmt.Println("=== 05: Modern Mathematics, PRNG, and Cryptography ===")

	// 1. Modern PRNG (Math/Rand seeded with local source, avoiding deprecated global Seed)
	source := mathRand.NewSource(time.Now().UnixNano())
	rng := mathRand.New(source)
	diceRoll := rng.Intn(6) + 1
	fmt.Printf("PRNG Local Dice Roll: %d\n", diceRoll)

	// 2. Cryptographically Secure Random Number Generation (CSPRNG)
	cryptoVal, err := GenerateCryptoInt(100)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("CSPRNG Secure Integer [0, 100): %d\n", cryptoVal)
	}

	token, err := GenerateSecureToken(16)
	if err != nil {
		fmt.Printf("Error generating token: %v\n", err)
	} else {
		fmt.Printf("CSPRNG 32-Char Hex Token: %s\n\n", token)
	}

	// 3. Floating-Point Inaccuracy vs Big.Float
	fmt.Println("--- Floating Point Drift vs Precision ---")
	standardFloatSum := 0.1 + 0.2
	fmt.Printf("Standard Float64 (0.1 + 0.2): %.17f\n", standardFloatSum)
	fmt.Printf("Is exactly equal to 0.3?     : %t\n", standardFloatSum == 0.3)

	bigSum, err := FinancialPrecisionAdd("0.1", "0.2")
	if err == nil {
		fmt.Printf("math/big.Float Precision Sum: %s\n\n", bigSum.Text('f', 20))
	}

	// 4. Standard Math Library functions
	hypotenuse := math.Hypot(3.0, 4.0)
	fmt.Printf("Standard Math Hypot(3, 4)   : %.2f\n", hypotenuse)
	fmt.Printf("2 to the power of 10        : %.0f\n", math.Pow(2, 10))
}
