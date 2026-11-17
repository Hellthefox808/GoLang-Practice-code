// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"strings"
	"testing"
)

func TestClassifyHTTPStatus(t *testing.T) {
	tests := []struct {
		code     int
		expected string
	}{
		{200, "Success (2xx)"},
		{404, "Client Error (4xx)"},
		{500, "Server Error (5xx)"},
		{999, "Unknown Status Code"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			if got := ClassifyHTTPStatus(tt.code); got != tt.expected {
				t.Errorf("got %s, expected %s", got, tt.expected)
			}
		})
	}
}

func TestInspectDynamicType(t *testing.T) {
	if !strings.Contains(InspectDynamicType(42), "Integer: 42") {
		t.Error("expected integer type match")
	}
	if !strings.Contains(InspectDynamicType("test"), "String: \"test\"") {
		t.Error("expected string type match")
	}
	if !strings.Contains(InspectDynamicType(true), "Boolean flag: true") {
		t.Error("expected boolean type match")
	}
}
