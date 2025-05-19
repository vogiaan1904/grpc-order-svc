package util

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrWeakPassword = errors.New("password not strong enough")
	ErrInvalidEmail = errors.New("invalid email format")
)

// IsValidEmail validates email format using a simple regex pattern
func IsValidEmail(email string) bool {
	email = strings.TrimSpace(email)

	// Simple email validation regex
	// For production, consider using a more comprehensive solution
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(pattern, email)
	return match
}

func IsStrongPassword(password string) bool {
	// Check minimum length (8 characters)
	if len(password) < 8 {
		return false
	}

	// Check for at least one uppercase letter
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	if !hasUpper {
		return false
	}

	// Check for at least one lowercase letter
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	if !hasLower {
		return false
	}

	// Check for at least one digit
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	if !hasDigit {
		return false
	}

	// Check for at least one special character
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`).MatchString(password)

	return hasSpecial
}
