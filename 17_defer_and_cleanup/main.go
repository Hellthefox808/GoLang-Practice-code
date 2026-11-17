// Package main details defer statements, LIFO stack mechanics, argument evaluation timing,
// named return modification, and graceful panic-and-recover error boundaries.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"fmt"
	"time"
)

// MeasureExecutionTime demonstrates defer closure pattern for automated function profiling.
func MeasureExecutionTime(operationName string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("[PROFILER] %s executed in %v\n", operationName, time.Since(start))
	}
}

// ModifyReturnViaDefer demonstrates that a deferred closure can read and alter named returns.
func ModifyReturnViaDefer() (result int) {
	defer func() {
		result *= 2 // Doubles the returned value before returning to caller!
	}()
	return 25 // Set to 25, then multiplied to 50 by deferred function
}

// SafeWorker demonstrates panic recovery to prevent an entire process or HTTP server crash.
func SafeWorker(taskID int, shouldPanic bool) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("task %d panicked and was recovered safely: %v", taskID, r)
		}
	}()

	if shouldPanic {
		panic("nil pointer dereference or fatal hardware fault simulated")
	}

	fmt.Printf("Task %d completed normally.\n", taskID)
	return nil
}

func main() {
	fmt.Println("=== 17: Defer Mechanics, LIFO Execution, and Panic/Recover ===")

	// 1. Deferred profiler pattern
	stopProfiler := MeasureExecutionTime("MainExecutionFlow")
	defer stopProfiler()

	// 2. LIFO (Last-In, First-Out) demonstration
	fmt.Println("\nScheduling deferred calls (Notice LIFO reversal on function exit):")
	for i := 1; i <= 3; i++ {
		// Notice: i is evaluated immediately as an argument
		defer fmt.Printf("  Deferred LIFO Call [%d]\n", i)
	}

	// 3. Named return mutation
	res := ModifyReturnViaDefer()
	fmt.Printf("\nNamed Return Value after Defer Mutation: %d (Expected: 50)\n\n", res)

	// 4. Panic & Recover Boundary
	fmt.Println("--- Testing Panic Recovery Boundary ---")
	err := SafeWorker(101, false)
	fmt.Printf("Worker 101 Result: err=%v\n", err)

	err = SafeWorker(102, true) // Will panic and be recovered
	fmt.Printf("Worker 102 Result: err=%v\n", err)
	fmt.Println("Main thread continues execution gracefully!")
}
