// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"testing"
	"time"
)

func TestParseRFC3339(t *testing.T) {
	iso := "2026-11-17T13:00:00Z"
	parsed, err := ParseRFC3339(iso)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if parsed.Year() != 2026 || parsed.Month() != time.November || parsed.Day() != 17 {
		t.Errorf("unexpected date components: %v", parsed)
	}

	_, err = ParseRFC3339("invalid-date-string")
	if err == nil {
		t.Error("expected error for malformed string, got nil")
	}
}

func TestFormatTimestamp(t *testing.T) {
	fixedTime := time.Date(2026, 11, 17, 10, 30, 45, 0, time.UTC)
	formatted := FormatTimestamp(fixedTime, StandardDateLayout)
	if formatted != "2026-11-17" {
		t.Errorf("expected 2026-11-17, got %s", formatted)
	}
}
