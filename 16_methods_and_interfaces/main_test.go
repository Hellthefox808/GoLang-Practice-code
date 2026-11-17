// Author: Ravi Ranjan Singh
// License: MIT

package main

import (
	"errors"
	"testing"
)

func TestUserAccountNotify(t *testing.T) {
	u := &UserAccount{Name: "Ravi", Email: "ravi@test.com", Notification: 0}
	err := u.Notify("Test alert")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if u.Notification != 1 {
		t.Errorf("expected notification count 1, got %d", u.Notification)
	}

	err = u.Notify("   ")
	if !errors.Is(err, ErrEmptyNotification) {
		t.Errorf("expected ErrEmptyNotification, got %v", err)
	}
}

func TestBroadcastAlert(t *testing.T) {
	u := &UserAccount{Name: "Ravi", Email: "ravi@test.com"}
	s := SlackChannel{WebhookURL: "https://slack.com"}

	channels := []Notifier{u, s}
	err := BroadcastAlert(channels, "Global broadcast")
	if err != nil {
		t.Errorf("broadcast failed: %v", err)
	}
}
