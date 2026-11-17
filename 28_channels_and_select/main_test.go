// Author: Ravi Ranjan Singh
// License: MIT

package main

import "testing"

func TestRunWorkerPool(t *testing.T) {
	jobs := []Job{
		{ID: 1, WorkData: "A"},
		{ID: 2, WorkData: "B"},
		{ID: 3, WorkData: "C"},
		{ID: 4, WorkData: "D"},
	}

	results := RunWorkerPool(2, jobs)

	if len(results) != len(jobs) {
		t.Fatalf("expected %d results, got %d", len(jobs), len(results))
	}

	seen := make(map[int]bool)
	for _, r := range results {
		seen[r.JobID] = true
	}

	for _, j := range jobs {
		if !seen[j.ID] {
			t.Errorf("job %d was not processed in results", j.ID)
		}
	}
}
