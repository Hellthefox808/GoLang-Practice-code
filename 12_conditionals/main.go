// Package main demonstrates idiomatic Go branching, the short-statement if idiom,
// early returns (guard clauses), and boolean logic short-circuiting.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"errors"
	"fmt"
)

var (
	ErrAccessDenied    = errors.New("access denied: insufficient permissions")
	ErrAccountInactive = errors.New("access denied: account is suspended")
)

type Account struct {
	Username  string
	Role      string
	IsActive  bool
	LoginRate int
}

// AuthorizeRequest evaluates access using the "Guard Clause" early return pattern,
// keeping the happy path non-nested and clean.
func AuthorizeRequest(acc Account, requiredRole string) (bool, error) {
	// Guard clause 1: Inactive account check
	if !acc.IsActive {
		return false, ErrAccountInactive
	}

	// Guard clause 2: Rate limit threshold check
	if acc.LoginRate > 100 {
		return false, fmt.Errorf("rate limit exceeded (%d requests/sec)", acc.LoginRate)
	}

	// Guard clause 3: Role comparison with short variable initialization
	if match := acc.Role == requiredRole || acc.Role == "admin"; !match {
		return false, ErrAccessDenied
	}

	// Happy Path is reached without deep nesting
	return true, nil
}

func main() {
	fmt.Println("=== 12: Conditionals, Guard Clauses, and Scoped If Statements ===")

	adminAccount := Account{
		Username:  "ravi",
		Role:      "admin",
		IsActive:  true,
		LoginRate: 25,
	}

	// 1. If statement with short variable initialization
	// 'authorized' and 'err' are strictly scoped to the if-else block
	if authorized, err := AuthorizeRequest(adminAccount, "moderator"); err != nil {
		fmt.Printf("Authorization Failed: %v\n", err)
	} else {
		fmt.Printf("Authorization Successful: Account '%s' authorized=%t\n", adminAccount.Username, authorized)
	}

	// 2. Testing rejected account
	suspendedAccount := Account{
		Username:  "intruder",
		Role:      "guest",
		IsActive:  false,
		LoginRate: 50,
	}

	if _, err := AuthorizeRequest(suspendedAccount, "user"); err != nil {
		fmt.Printf("Defense check working: %v\n", err)
	}
}
