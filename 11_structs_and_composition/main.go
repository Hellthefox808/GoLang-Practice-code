// Package main demonstrates idiomatic object modeling in Go through structs,
// composition over inheritance via struct embedding, field promotion, constructors,
// and struct memory alignment optimizations.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"unsafe"
)

var (
	ErrInvalidEmail = errors.New("email address must contain '@'")
	ErrUnderageUser = errors.New("user age must be at least 18")
)

// BaseEntity holds audit metadata, intended for struct embedding across domain entities.
type BaseEntity struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// User demonstrates struct embedding (composition over inheritance).
// Notice field ordering: int64/time.Time (8 bytes), string (16 bytes), bool (1 byte)
// Grouping larger primitive types first minimizes struct memory padding.
type User struct {
	BaseEntity        // Embedded struct: its fields and methods are promoted to User directly!
	Email      string `json:"email"`
	FullName   string `json:"full_name"`
	Age        int    `json:"age"`
	IsActive   bool   `json:"is_active"`
}

// UnoptimizedStruct demonstrates memory padding waste due to poor field ordering.
type UnoptimizedStruct struct {
	A bool  // 1 byte + 7 bytes padding
	B int64 // 8 bytes
	C bool  // 1 byte + 7 bytes padding
} // Total: 24 bytes

// OptimizedStruct groups fields to eliminate alignment padding.
type OptimizedStruct struct {
	B int64 // 8 bytes
	A bool  // 1 byte
	C bool  // 1 byte + 6 bytes padding
} // Total: 16 bytes (33% memory saved!)

// NewUser acts as an idiomatic constructor, enforcing validation invariant boundaries.
func NewUser(id, fullName, email string, age int) (*User, error) {
	if !strings.Contains(email, "@") {
		return nil, ErrInvalidEmail
	}
	if age < 18 {
		return nil, ErrUnderageUser
	}

	now := time.Now().UTC()
	return &User{
		BaseEntity: BaseEntity{
			ID:        id,
			CreatedAt: now,
			UpdatedAt: now,
		},
		FullName: fullName,
		Email:    email,
		Age:      age,
		IsActive: true,
	}, nil
}

func main() {
	fmt.Println("=== 11: Structs, Composition, and Memory Alignment ===")

	// 1. Constructor instantiation with invariant validation
	user, err := NewUser("usr_1001", "Ravi Ranjan Singh", "ravi@example.com", 25)
	if err != nil {
		fmt.Printf("User creation error: %v\n", err)
		return
	}

	// 2. Direct field promotion access
	// Notice we can access user.ID directly without writing user.BaseEntity.ID
	fmt.Printf("User ID (Promoted) : %s\n", user.ID)
	fmt.Printf("User Full Name     : %s\n", user.FullName)
	fmt.Printf("User Email         : %s\n", user.Email)
	fmt.Printf("User CreatedAt     : %s\n", user.CreatedAt.Format(time.RFC3339))
	fmt.Printf("Full Struct Detail : %+v\n\n", user)

	// 3. Validation rejection test
	_, err = NewUser("usr_1002", "Junior Dev", "invalid-email", 16)
	if err != nil {
		fmt.Printf("Validation caught invalid user correctly: %v\n\n", err)
	}

	// 4. Memory Layout & Struct Alignment Optimization Demonstration
	fmt.Println("--- Struct Memory Alignment Optimization ---")
	fmt.Printf("Size of UnoptimizedStruct: %d bytes\n", unsafe.Sizeof(UnoptimizedStruct{}))
	fmt.Printf("Size of OptimizedStruct  : %d bytes\n", unsafe.Sizeof(OptimizedStruct{}))
}
