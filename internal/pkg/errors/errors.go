package errors

import "fmt"

// ErrorCode represents a custom error code.
type ErrorCode string

// Define custom error codes.
const (
	InvalidArgumentError ErrorCode = "InvalidArgument"
	NotFoundError        ErrorCode = "NotFound"
	InternalError        ErrorCode = "Internal"
	// Add more error codes as needed...
)

// Error represents a custom error.
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// Error returns the error message.
func (e *Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewInvalidArgumentError creates a new invalid argument error.
func NewInvalidArgumentError(message string) *Error {
	return &Error{
		Code:    InvalidArgumentError,
		Message: message,
	}
}

// NewNotFoundError creates a new not found error.
func NewNotFoundError(message string) *Error {
	return &Error{
		Code:    NotFoundError,
		Message: message,
	}
}

// NewInternalError creates a new internal error.
func NewInternalError(message string) *Error {
	return &Error{
		Code:    InternalError,
		Message: message,
	}
}

// Add more custom error creation functions as needed...
