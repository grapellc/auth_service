package errors

import "fmt"

// DomainError represents a domain-specific error
type DomainError struct {
	Code    string
	Message string
	Details map[string]interface{}
}

func (e *DomainError) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewDomainError creates a new domain error
func NewDomainError(code, message string, details map[string]interface{}) *DomainError {
	return &DomainError{
		Code:    code,
		Message: message,
		Details: details,
	}
}

// Predefined domain errors
var (
	ErrProductNotFound     = NewDomainError("PRODUCT_NOT_FOUND", "Product not found", nil)
	ErrProductInvalidName  = NewDomainError("PRODUCT_INVALID_NAME", "Product name is invalid", nil)
	ErrProductInvalidPrice = NewDomainError("PRODUCT_INVALID_PRICE", "Product price is invalid", nil)
	ErrUserNotFound        = NewDomainError("USER_NOT_FOUND", "User not found", nil)
	ErrUserInvalidEmail    = NewDomainError("USER_INVALID_EMAIL", "User email is invalid", nil)
)
