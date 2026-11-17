// Package main explores hash table architecture in Go maps, the comma-ok idiom,
// randomized iteration orders, delete semantics, nil map hazards, and the modern maps package.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"fmt"
	"maps"
)

// LookupLanguage retrieves a language full name using the comma-ok idiom to distinguish
// between a missing key and a key holding a zero-value.
func LookupLanguage(registry map[string]string, code string) (string, bool) {
	val, exists := registry[code]
	return val, exists
}

func main() {
	fmt.Println("=== 10: Hash Tables, Map Semantics, and Safety ===")

	// 1. Initializing with capacity hint
	// Providing a size hint avoids bucket reallocation and rehashing overhead
	languages := make(map[string]string, 8)
	languages["GO"] = "Golang (Go)"
	languages["RS"] = "Rust"
	languages["TS"] = "TypeScript"
	languages["PY"] = "Python"
	languages["CPP"] = "C++"

	fmt.Printf("Registry size: %d entries\n", len(languages))

	// 2. Comma-ok idiom (Essential to distinguish between zero-value vs missing key)
	if val, ok := LookupLanguage(languages, "GO"); ok {
		fmt.Printf("Lookup 'GO' found: %s\n", val)
	}

	if _, ok := LookupLanguage(languages, "JAVA"); !ok {
		fmt.Println("Lookup 'JAVA': Key does not exist in registry.")
	}

	// 3. Deletion semantics
	// delete() is idempotent: deleting a non-existent key is a safe no-op
	delete(languages, "PY")
	delete(languages, "NON_EXISTENT_KEY") // Safe no-op

	// 4. Map iteration is non-deterministic by language design (security against HashDOS attacks)
	fmt.Println("\nIterating over map entries (Order is randomized each run):")
	for code, name := range languages {
		fmt.Printf("  [%-3s] => %s\n", code, name)
	}

	// 5. Modern Go 1.21+ maps package features
	clonedRegistry := maps.Clone(languages)
	fmt.Printf("\nCloned Registry Equal? %t\n", maps.Equal(languages, clonedRegistry))

	// Delete entries meeting predicate
	maps.DeleteFunc(clonedRegistry, func(k, v string) bool {
		return k == "CPP" || k == "RS"
	})
	fmt.Printf("Cloned after DeleteFunc: %v\n", clonedRegistry)

	// 6. The Nil Map Trap
	// Reading from a nil map is safe (returns zero-value), but WRITING to a nil map PANICS!
	var nilMap map[string]int
	fmt.Printf("\nReading from nil map: %d (Safe zero-value)\n", nilMap["anyKey"])
	// nilMap["key"] = 100 // FATAL PANIC: panic: assignment to entry in nil map
	fmt.Println("Rule: Always initialize maps via make() or literal before writing!")
}
