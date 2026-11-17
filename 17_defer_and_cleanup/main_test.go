// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"strings"
	"testing"
)

func TestModifyReturnViaDefer(t *testing.T) {
	if got := ModifyReturnViaDefer(); got != 50 {
		t.Errorf("expected 50, got %d", got)
	}
}

func TestSafeWorker(t *testing.T) {
	err := SafeWorker(1, false)
	if err != nil {
		t.Errorf("expected nil error for normal run, got %v", err)
	}

	err = SafeWorker(2, true)
	if err == nil || !strings.Contains(err.Error(), "panicked and was recovered safely") {
		t.Errorf("expected panic recovery error, got %v", err)
	}
}
