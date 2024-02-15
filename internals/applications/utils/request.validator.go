package utils

import (
	"errors"
	"regexp"
)

// Validator represents a request validator
type Validator struct{}

// NewValidator creates a new instance of Validator
func NewValidator() *Validator {
	return &Validator{}
}

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
}

// ValidationErrors represents a list of validation errors
type ValidationErrors []ValidationError

// Add adds a new validation error to the list
func (ve *ValidationErrors) Add(field, message string) {
	*ve = append(*ve, ValidationError{Field: field, Message: message})
}

// ValidateString checks if a string value matches the given validation rules
func (v *Validator) ValidateString(value, fieldName string, rules ...string) error {
	for _, rule := range rules {
		switch rule {
		case "required":
			if value == "" {
				return errors.New(fieldName + " is required")
			}
		case "email":
			if !isValidEmail(value) {
				return errors.New(fieldName + " must be a valid email address")
			}
			// Add more validation rules as needed
		}
	}
	return nil
}

// isValidEmail checks if a string is a valid email address
func isValidEmail(email string) bool {
	// This is a simple regex pattern for email validation
	// You may use a more sophisticated pattern depending on your requirements
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
