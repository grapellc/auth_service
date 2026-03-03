package errors

import (
	"net/http"
)

// HTTPError represents an HTTP error
type HTTPError struct {
	StatusCode int
	Message    string
	Details    map[string]interface{}
}

func (e *HTTPError) Error() string {
	return e.Message
}

// NewHTTPError creates a new HTTP error
func NewHTTPError(statusCode int, message string, details map[string]interface{}) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
		Details:    details,
	}
}

// Predefined HTTP errors
var (
	ErrBadRequest          = NewHTTPError(http.StatusBadRequest, "Bad Request", nil)
	ErrUnauthorized        = NewHTTPError(http.StatusUnauthorized, "Unauthorized", nil)
	ErrForbidden           = NewHTTPError(http.StatusForbidden, "Forbidden", nil)
	ErrNotFound            = NewHTTPError(http.StatusNotFound, "Not Found", nil)
	ErrInternalServerError = NewHTTPError(http.StatusInternalServerError, "Internal Server Error", nil)
)
