// Package main details methods, value vs pointer receivers, implicit interface satisfaction,
// method sets, polymorphic dispatch, and the nil interface trap.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrEmptyNotification = errors.New("notification body cannot be empty")
)

// Notifier defines the behavioral contract for outbound alerting channels.
type Notifier interface {
	Notify(message string) error
	ChannelName() string
}

// UserAccount represents a core domain entity.
type UserAccount struct {
	Name         string
	Email        string
	Notification int
}

// Value Receiver: Operates on a copy of UserAccount.
// Calling this does NOT mutate the original caller instance!
func (u UserAccount) DisplayProfile() string {
	return fmt.Sprintf("User: %s <%s> [Notifications Received: %d]", u.Name, u.Email, u.Notification)
}

// Pointer Receiver: Operates on *UserAccount.
// Mutates the actual caller struct instance in place and satisfies Notifier interface!
func (u *UserAccount) Notify(message string) error {
	if u == nil {
		return errors.New("cannot notify nil user")
	}
	if strings.TrimSpace(message) == "" {
		return ErrEmptyNotification
	}
	u.Notification++
	fmt.Printf("[Email Alert -> %s]: %s\n", u.Email, message)
	return nil
}

func (u *UserAccount) ChannelName() string {
	return "Direct-Email"
}

// SlackChannel represents a webhook-based notifier.
type SlackChannel struct {
	WebhookURL string
}

func (s SlackChannel) Notify(message string) error {
	if strings.TrimSpace(message) == "" {
		return ErrEmptyNotification
	}
	fmt.Printf("[Slack Webhook -> %s]: %s\n", s.WebhookURL, message)
	return nil
}

func (s SlackChannel) ChannelName() string {
	return "Slack-Bot"
}

// BroadcastAlert broadcasts an alert across any collection of Notifier implementations.
func BroadcastAlert(notifiers []Notifier, message string) error {
	for _, n := range notifiers {
		if err := n.Notify(message); err != nil {
			return fmt.Errorf("dispatch failed on %s: %w", n.ChannelName(), err)
		}
	}
	return nil
}

func main() {
	fmt.Println("=== 16: Methods, Value vs Pointer Receivers, and Interfaces ===")

	// 1. Value Receiver vs Pointer Receiver Demonstration
	user := UserAccount{
		Name:         "Ravi Ranjan Singh",
		Email:        "ravi@example.com",
		Notification: 0,
	}

	fmt.Println("Initial Profile        :", user.DisplayProfile())

	// Call pointer receiver method
	_ = user.Notify("System upgrade scheduled for 02:00 UTC")
	fmt.Println("After Notify (Count++):", user.DisplayProfile())

	// 2. Polymorphism via Interface Contracts
	slack := SlackChannel{WebhookURL: "https://hooks.slack.com/services/T00/B00/X00"}

	// Notice: &user is passed because *UserAccount implements Notifier (pointer receiver)
	channels := []Notifier{&user, slack}

	fmt.Println("\n--- Dispatching Broadcast Notification ---")
	err := BroadcastAlert(channels, "Production deployment completed successfully.")
	if err != nil {
		fmt.Printf("Broadcast failed: %v\n", err)
	}

	// 3. Inspecting the interface at runtime
	for _, ch := range channels {
		fmt.Printf("Registered Channel: %-15s (Underlying type: %T)\n", ch.ChannelName(), ch)
	}
}
