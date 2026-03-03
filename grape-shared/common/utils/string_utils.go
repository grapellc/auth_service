package utils

import "strings"

// StringPtr returns a pointer to the given string
func StringPtr(s string) *string {
	return &s
}

// IsEmpty checks if a string is empty or contains only whitespace
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// Truncate truncates a string to the specified length
func Truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}
