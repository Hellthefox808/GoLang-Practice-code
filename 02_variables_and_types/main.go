// Package main explores the comprehensive Go static type system, declaration
// syntaxes, variable scopes, zero values, and constant definitions with iota.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"fmt"
	"math"
)

// Exported constant (PascalCase) accessible outside this package.
const AppVersion string = "2.4.0"

// Unexported constant (camelCase) scoped strictly to this package.
const defaultClusterRegion = "us-east-1"

// Iota enumerator demonstration for distinct typed constants.
type Role uint8

const (
	RoleGuest Role = iota
	RoleUser
	RoleModerator
	RoleAdministrator
)

func (r Role) String() string {
	switch r {
	case RoleGuest:
		return "Guest"
	case RoleUser:
		return "User"
	case RoleModerator:
		return "Moderator"
	case RoleAdministrator:
		return "Administrator"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Println("=== 02: Variables, Static Types, and Zero Values ===")

	// 1. Explicit variable declarations
	var developerName string = "Ravi Ranjan Singh"
	var isActiveDeveloper bool = true
	var activeGoroutinesCount uint16 = 1024
	var piApproximation float64 = 3.141592653589793

	fmt.Printf("Explicit String   : %-20s | Type: %T\n", developerName, developerName)
	fmt.Printf("Explicit Boolean  : %-20t | Type: %T\n", isActiveDeveloper, isActiveDeveloper)
	fmt.Printf("Explicit Uint16   : %-20d | Type: %T\n", activeGoroutinesCount, activeGoroutinesCount)
	fmt.Printf("Explicit Float64  : %-20.10f | Type: %T\n\n", piApproximation, piApproximation)

	// 2. Zero Values in Go (Go guarantees memory safety with deterministic defaults)
	var zeroString string
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroPointer *int

	fmt.Println("--- Deterministic Zero Values ---")
	fmt.Printf("string default : %q\n", zeroString)
	fmt.Printf("int default    : %d\n", zeroInt)
	fmt.Printf("float default  : %f\n", zeroFloat)
	fmt.Printf("bool default   : %t\n", zeroBool)
	fmt.Printf("pointer default: %v\n\n", zeroPointer)

	// 3. Type inference with 'var'
	var serviceEndpoint = "https://api.github.com/v1"
	fmt.Printf("Inferred String   : %s (Type: %T)\n", serviceEndpoint, serviceEndpoint)

	// 4. Short variable declaration (:=) allowed only within function scopes
	connectionTimeoutSec := 30
	systemLoad := 0.85
	fmt.Printf("Short Declared    : Timeout=%d, Load=%.2f\n\n", connectionTimeoutSec, systemLoad)

	// 5. Numeric boundary exploration
	fmt.Println("--- Standard Numeric Limits ---")
	fmt.Printf("MaxInt8  : %d | MinInt8  : %d\n", math.MaxInt8, math.MinInt8)
	fmt.Printf("MaxUint8 : %d\n", math.MaxUint8)
	fmt.Printf("MaxInt64 : %d\n", math.MaxInt64)
	fmt.Printf("MaxFloat64: %e\n\n", math.MaxFloat64)

	// 6. Iota Enum inspection
	fmt.Println("--- Iota Enumerations ---")
	fmt.Printf("RoleGuest         : %d (%s)\n", RoleGuest, RoleGuest)
	fmt.Printf("RoleAdministrator : %d (%s)\n", RoleAdministrator, RoleAdministrator)
	fmt.Printf("App Version       : %s\n", AppVersion)
}
