package errors

import "fmt"

// ValidationError represents a validation error
type ValidationError struct {
	Field   string
	Message string
	Value   interface{}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation failed for field '%s': %s", e.Field, e.Message)
}

// ValidationErrors represents multiple validation errors
type ValidationErrors struct {
	Errors []ValidationError
}

func (e *ValidationErrors) Error() string {
	if len(e.Errors) == 0 {
		return "no validation errors"
	}
	return fmt.Sprintf("validation failed: %d errors", len(e.Errors))
}

func (e *ValidationErrors) Add(field, message string, value interface{}) {
	e.Errors = append(e.Errors, ValidationError{
		Field:   field,
		Message: message,
		Value:   value,
	})
}

func (e *ValidationErrors) HasErrors() bool {
	return len(e.Errors) > 0
}
