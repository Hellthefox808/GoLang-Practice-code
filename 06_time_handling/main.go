// Package main details time representation, the Go reference layout idiom,
// ISO8601/RFC3339 serialization, timezone conversions, monotonic clock measurements,
// and time.Ticker/Timer primitives.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"fmt"
	"time"
)

// The canonical Go reference date: Mon Jan 2 15:04:05 MST 2006 (01/02 03:04:05PM '06 -0700)
const (
	StandardDateLayout     = "2006-01-02"
	StandardTimeLayout     = "15:04:05"
	HumanReadableFull      = "Monday, 02-Jan-2006 15:04:05 MST"
	MicroserviceAuditStamp = "2006/01/02 15:04:05.000000"
)

// FormatTimestamp transforms a time.Time instance into a specified layout.
func FormatTimestamp(t time.Time, layout string) string {
	return t.Format(layout)
}

// ParseRFC3339 parses an RFC3339 string securely into a UTC time.Time.
func ParseRFC3339(isoString string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, isoString)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse RFC3339 timestamp %q: %w", isoString, err)
	}
	return t.UTC(), nil
}

// ConvertToLocation shifts a time instance to the target IANA timezone location.
func ConvertToLocation(t time.Time, locationName string) (time.Time, error) {
	loc, err := time.LoadLocation(locationName)
	if err != nil {
		return time.Time{}, fmt.Errorf("unknown timezone location %q: %w", locationName, err)
	}
	return t.In(loc), nil
}

func main() {
	fmt.Println("=== 06: Time Engineering & Monotonic Measurements ===")

	// 1. Current Time & Formatting
	now := time.Now()
	fmt.Println("Current Wall/Monotonic Time:", now)
	fmt.Println("Formatted (Human Readable) :", FormatTimestamp(now, HumanReadableFull))
	fmt.Println("Formatted (Audit Stamp)    :", FormatTimestamp(now, MicroserviceAuditStamp))
	fmt.Println("Formatted (Standard RFC3339):", FormatTimestamp(now, time.RFC3339Nano))

	// 2. Specific Date Construction
	projectReleaseDate := time.Date(2026, time.November, 17, 13, 0, 0, 0, time.UTC)
	fmt.Println("\nTarget Release Date (UTC)  :", projectReleaseDate.Format(time.RFC3339))

	// 3. Time Parsing
	sampleISO := "2026-12-31T23:59:59Z"
	parsedTime, err := ParseRFC3339(sampleISO)
	if err != nil {
		fmt.Printf("Parsing failed: %v\n", err)
	} else {
		fmt.Printf("Parsed UTC Instance        : %v (Year: %d, Day: %d)\n",
			parsedTime, parsedTime.Year(), parsedTime.YearDay())
	}

	// 4. Time Zone Conversion (IANA)
	tokyoTime, err := ConvertToLocation(now, "Asia/Tokyo")
	if err == nil {
		fmt.Println("Current Time in Tokyo      :", tokyoTime.Format(HumanReadableFull))
	}

	istTime, err := ConvertToLocation(now, "Asia/Kolkata")
	if err == nil {
		fmt.Println("Current Time in New Delhi  :", istTime.Format(HumanReadableFull))
	}

	// 5. Elapsed Duration & Benchmarking using Monotonic Clock
	start := time.Now()
	// Simulate lightweight operation
	sum := 0
	for i := 0; i < 1_000_000; i++ {
		sum += i
	}
	elapsed := time.Since(start)
	fmt.Printf("\nElapsed time (Monotonic)   : %v (Sum check: %d)\n", elapsed, sum)

	// 6. Ticker demonstration
	fmt.Println("\nRunning rapid 2-tick demonstration:")
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	ticksReceived := 0
	for tick := range ticker.C {
		fmt.Printf("  Tick [%d] received at %s\n", ticksReceived+1, tick.Format(StandardTimeLayout))
		ticksReceived++
		if ticksReceived >= 2 {
			break
		}
	}
}
