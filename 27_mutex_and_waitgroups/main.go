// Package main demonstrates mutual exclusion with sync.Mutex, sync.RWMutex,
// lock-free atomic primitives with sync/atomic, and race condition prevention.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ThreadSafeMetrics protects shared metrics using sync.RWMutex.
type ThreadSafeMetrics struct {
	mu     sync.RWMutex
	counts map[string]int
}

func NewThreadSafeMetrics() *ThreadSafeMetrics {
	return &ThreadSafeMetrics{
		counts: make(map[string]int),
	}
}

// Increment acquires an exclusive write lock.
func (m *ThreadSafeMetrics) Increment(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.counts[key]++
}

// GetCount acquires a shared read lock, allowing thousands of simultaneous readers.
func (m *ThreadSafeMetrics) GetCount(key string) int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.counts[key]
}

func main() {
	fmt.Println("=== 27: Mutual Exclusion, RWMutex, and Lock-Free Atomics ===")

	metrics := NewThreadSafeMetrics()
	var wg sync.WaitGroup

	// 1. Spawning 50 concurrent writers
	numWriters := 50
	for i := 0; i < numWriters; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			metrics.Increment("api_requests")
		}(i)
	}

	// 2. Spawning 20 concurrent readers simultaneously
	numReaders := 20
	for i := 0; i < numReaders; i++ {
		wg.Add(1)
		go func(readerID int) {
			defer wg.Done()
			_ = metrics.GetCount("api_requests")
		}(i)
	}

	wg.Wait()
	fmt.Printf("RWMutex Synchronized Value: 'api_requests' = %d (Expected: %d)\n\n",
		metrics.GetCount("api_requests"), numWriters)

	// 3. Lock-Free Atomic Operations (sync/atomic)
	// For simple counters and flags, sync/atomic operates via direct CPU instructions (e.g. LOCK CMPXCHG)
	// without OS thread context switches, providing immense performance gains.
	var atomicCounter atomic.Int64
	var atomicWG sync.WaitGroup

	operationsCount := 1000
	start := time.Now()

	for i := 0; i < operationsCount; i++ {
		atomicWG.Add(1)
		go func() {
			defer atomicWG.Done()
			atomicCounter.Add(1)
		}()
	}

	atomicWG.Wait()
	elapsed := time.Since(start)

	fmt.Println("--- Lock-Free Atomic Counter (sync/atomic.Int64) ---")
	fmt.Printf("Final Counter Value : %d\n", atomicCounter.Load())
	fmt.Printf("Execution Time      : %v (Across %d concurrent goroutines)\n", elapsed, operationsCount)
}
