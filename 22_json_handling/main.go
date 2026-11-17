// Package main demonstrates enterprise JSON serialization, struct tags, streaming decoders,
// DisallowUnknownFields security guards, custom Marshal/Unmarshal interfaces, and polymorphic json.RawMessage.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"
)

// Course represents an educational asset with rich JSON tags.
type Course struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	PriceUSD  float64   `json:"price_usd"`
	Tags      []string  `json:"tags,omitempty"`  // Omitted if slice is empty/nil
	SecretKey string    `json:"-"`               // Completely hidden from JSON serialization
	CreatedAt time.Time `json:"created_at"`
}

// StrictCoursePayload validates incoming JSON payloads strictly without allowing unexpected fields.
func DecodeStrictJSON(payload []byte) (*Course, error) {
	var c Course
	decoder := json.NewDecoder(bytes.NewReader(payload))
	decoder.DisallowUnknownFields() // Security: Rejects unexpected payload keys!

	if err := decoder.Decode(&c); err != nil {
		return nil, fmt.Errorf("strict decode failed: %w", err)
	}

	if c.Title == "" {
		return nil, errors.New("course title cannot be empty")
	}

	return &c, nil
}

// EventEnvelope demonstrates polymorphic JSON handling via json.RawMessage.
type EventEnvelope struct {
	EventType string          `json:"event_type"`
	Timestamp int64           `json:"timestamp"`
	Payload   json.RawMessage `json:"payload"` // Deferred parsing based on event_type
}

type UserRegisteredEvent struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	fmt.Println("=== 22: Advanced JSON Serialization, Tags, and Strict Decoding ===")

	// 1. Serialization (MarshalIndent)
	courses := []Course{
		{
			ID:        "crs_101",
			Title:     "Advanced Go Concurrency",
			PriceUSD:  49.99,
			Tags:      []string{"golang", "goroutines", "systems"},
			SecretKey: "super_secret_internal_key_never_leak",
			CreatedAt: time.Now().UTC(),
		},
		{
			ID:        "crs_102",
			Title:     "Kubernetes Orchestration in Go",
			PriceUSD:  59.99,
			Tags:      nil, // Will be omitted by 'omitempty'
			SecretKey: "internal_db_salt",
			CreatedAt: time.Now().UTC(),
		},
	}

	encodedJSON, err := json.MarshalIndent(courses, "", "  ")
	if err != nil {
		fmt.Printf("Serialization error: %v\n", err)
		return
	}
	fmt.Println("--- Marshaled JSON Output (Notice 'tags' omitted in #2, 'SecretKey' completely absent) ---")
	fmt.Println(string(encodedJSON))

	// 2. Strict Decoding with DisallowUnknownFields
	validJSON := []byte(`{
		"id": "crs_200",
		"title": "Cloud-Native Microservices",
		"price_usd": 79.50,
		"created_at": "2026-11-17T13:00:00Z"
	}`)

	courseObj, err := DecodeStrictJSON(validJSON)
	if err != nil {
		fmt.Printf("Decode error: %v\n", err)
	} else {
		fmt.Printf("\nStrict Decoded Object: %+v\n", courseObj)
	}

	// Unknown field attack attempt
	maliciousJSON := []byte(`{
		"id": "crs_201",
		"title": "Exploit Test",
		"price_usd": 0.0,
		"created_at": "2026-11-17T13:00:00Z",
		"unexpected_malicious_key": "injection_payload"
	}`)

	_, err = DecodeStrictJSON(maliciousJSON)
	if err != nil {
		fmt.Printf("\nDisallowUnknownFields correctly rejected invalid payload:\n  %v\n\n", err)
	}

	// 3. Polymorphic Event Processing via json.RawMessage
	eventJSON := []byte(`{
		"event_type": "user_registered",
		"timestamp": 1794918000,
		"payload": {
			"username": "raviranjansingh",
			"email": "ravi@example.com"
		}
	}`)

	var envelope EventEnvelope
	_ = json.Unmarshal(eventJSON, &envelope)
	fmt.Println("--- Polymorphic Event Processing ---")
	fmt.Printf("Envelope Event Type: %s\n", envelope.EventType)

	if envelope.EventType == "user_registered" {
		var userEvt UserRegisteredEvent
		_ = json.Unmarshal(envelope.Payload, &userEvt)
		fmt.Printf("Extracted Concrete Event: User=%s, Email=%s\n", userEvt.Username, userEvt.Email)
	}
}
