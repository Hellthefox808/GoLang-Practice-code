// Package main explores URL RFC 3986 parsing, query parameter encoding,
// relative reference resolution, and programmatic url.URL struct construction.
//
// Author: Ravi Ranjan Singh
// Repository: https://github.com/Hellthefox808/GoLang-Practice-code
// License: MIT

package main

import (
	"fmt"
	"net/url"
)

// ParseAndInspectURL parses a raw URL string into its discrete RFC 3986 components.
func ParseAndInspectURL(rawURL string) (*url.URL, url.Values, error) {
	parsed, err := url.Parse(rawURL)
	if err != nil {
		return nil, nil, fmt.Errorf("invalid URL structure: %w", err)
	}

	queryParams := parsed.Query()
	return parsed, queryParams, nil
}

// BuildAPIEndpoint constructs an escaped URL safely without string concatenation vulnerabilities.
func BuildAPIEndpoint(baseURL, path string, params map[string]string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base url: %w", err)
	}

	rel, err := url.Parse(path)
	if err != nil {
		return "", fmt.Errorf("invalid relative path: %w", err)
	}

	resolved := base.ResolveReference(rel)

	query := resolved.Query()
	for k, v := range params {
		query.Set(k, v)
	}
	resolved.RawQuery = query.Encode()

	return resolved.String(), nil
}

func main() {
	fmt.Println("=== 20: RFC 3986 URL Parsing & Safe Query Manipulation ===")

	sampleURL := "https://api.github.com:443/repos/Hellthefox808/GoLang-Practice-code/issues?state=closed&per_page=50#readme"
	fmt.Printf("Analyzing raw URL:\n%s\n\n", sampleURL)

	parsed, query, err := ParseAndInspectURL(sampleURL)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// 1. Component Extraction
	fmt.Println("--- URL Component Breakdown ---")
	fmt.Printf("Scheme   : %s\n", parsed.Scheme)
	fmt.Printf("Host     : %s\n", parsed.Host)
	fmt.Printf("Hostname : %s\n", parsed.Hostname())
	fmt.Printf("Port     : %s\n", parsed.Port())
	fmt.Printf("Path     : %s\n", parsed.Path)
	fmt.Printf("RawQuery : %s\n", parsed.RawQuery)
	fmt.Printf("Fragment : %s\n\n", parsed.Fragment)

	// 2. Query Parameter inspection
	fmt.Println("--- Query Parameters (url.Values) ---")
	for key, values := range query {
		fmt.Printf("  Key [%-10s] => %v\n", key, values)
	}

	// 3. Programmatic URL Construction (Resolving the classic RawPath vs RawQuery bug)
	// IMPORTANT GO MECHANIC:
	// RawPath is NOT for query strings! RawPath is solely for preserving percent-escaped paths.
	// RawQuery is the correct field for query strings (e.g., "key=val&foo=bar").
	builtURL := &url.URL{
		Scheme:   "https",
		Host:     "api.example.com",
		Path:     "/v1/users",
		RawQuery: "developer=ravi_ranjan_singh&role=architect",
	}
	fmt.Printf("\nProgrammatically Built URL: %s\n", builtURL.String())

	// 4. Safe URL Builder with Relative Resolution
	endpoint, err := BuildAPIEndpoint("https://cdn.example.com/assets/", "icons/go.svg", map[string]string{
		"theme": "dark",
		"size":  "64",
	})
	if err == nil {
		fmt.Printf("Resolved Endpoint         : %s\n", endpoint)
	}
}
