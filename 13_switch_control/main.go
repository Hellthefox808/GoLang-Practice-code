// Package main explores idiomatic switch statements in Go: expression switches,
// implicit breaks, tagless switches as clean boolean ladders, and type switches for runtime interface inspection.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ClassifyHTTPStatus uses a standard expression switch with multiple matches per case.
func ClassifyHTTPStatus(statusCode int) string {
	switch statusCode {
	case 200, 201, 204:
		return "Success (2xx)"
	case 301, 302, 304:
		return "Redirection (3xx)"
	case 400, 401, 403, 404, 429:
		return "Client Error (4xx)"
	case 500, 502, 503, 504:
		return "Server Error (5xx)"
	default:
		return "Unknown Status Code"
	}
}

// EvaluatePerformance demonstrates a tagless switch acting as a clean boolean ladder.
func EvaluatePerformance(latencyMs int) string {
	switch {
	case latencyMs < 50:
		return "Ultra Low Latency (Excellent)"
	case latencyMs < 200:
		return "Normal Latency (Acceptable)"
	case latencyMs < 1000:
		return "Degraded Latency (Investigate)"
	default:
		return "Critical Latency Breach"
	}
}

// InspectDynamicType uses a type switch to inspect the concrete type behind an any interface.
func InspectDynamicType(val any) string {
	switch concrete := val.(type) {
	case int:
		return fmt.Sprintf("Integer: %d (doubled=%d)", concrete, concrete*2)
	case string:
		return fmt.Sprintf("String: %q (len=%d)", concrete, len(concrete))
	case bool:
		return fmt.Sprintf("Boolean flag: %t", concrete)
	case []string:
		return fmt.Sprintf("Slice of strings with %d elements", len(concrete))
	default:
		return fmt.Sprintf("Unsupported type: %T", concrete)
	}
}

func main() {
	fmt.Println("=== 13: Switch Statements, Tagless Ladders, and Type Switches ===")

	// 1. Expression switch with deterministic RNG
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	diceRoll := rng.Intn(6) + 1
	fmt.Printf("Rolled a %d: ", diceRoll)

	switch diceRoll {
	case 1:
		fmt.Println("Initial turn, spawn player.")
	case 2, 3, 4:
		fmt.Printf("Move forward %d spaces.\n", diceRoll)
	case 5:
		fmt.Println("Bonus square unlocked.")
	case 6:
		fmt.Println("Maximum roll! Roll again.")
	}

	// 2. HTTP Status classification
	fmt.Printf("\nStatus 201: %s\n", ClassifyHTTPStatus(201))
	fmt.Printf("Status 404: %s\n", ClassifyHTTPStatus(404))
	fmt.Printf("Status 503: %s\n", ClassifyHTTPStatus(503))

	// 3. Tagless switch
	fmt.Printf("\nLatency 35ms  : %s\n", EvaluatePerformance(35))
	fmt.Printf("Latency 1200ms: %s\n", EvaluatePerformance(1200))

	// 4. Type switch
	fmt.Println("\n--- Type Switch Inspection ---")
	fmt.Println(InspectDynamicType(1024))
	fmt.Println(InspectDynamicType("Ravi Ranjan Singh"))
	fmt.Println(InspectDynamicType([]string{"Go", "Docker", "K8s"}))
	fmt.Println(InspectDynamicType(3.14159))
}
