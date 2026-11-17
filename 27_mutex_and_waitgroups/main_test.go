// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"sync"
	"testing"
)

func TestThreadSafeMetrics(t *testing.T) {
	metrics := NewThreadSafeMetrics()
	var wg sync.WaitGroup

	iterations := 100
	for i := 0; i < iterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			metrics.Increment("hits")
		}()
	}

	wg.Wait()

	if val := metrics.GetCount("hits"); val != iterations {
		t.Errorf("expected %d hits, got %d", iterations, val)
	}
}
