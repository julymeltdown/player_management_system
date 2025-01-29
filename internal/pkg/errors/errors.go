package errors

import "fmt"

// ErrorCode represents a custom error code.
type ErrorCode string

// Define custom error codes.
const (
	InvalidArgumentError ErrorCode = "InvalidArgument"
	NotFoundError        ErrorCode = "NotFound"
	InternalError        ErrorCode = "Internal"
	DatabaseError        ErrorCode = "DatabaseError"
	NotConnectedError    ErrorCode = "NotConnected"
)

// Error messages for each error code.
var errorMessages = map[ErrorCode]string{
	InvalidArgumentError: "Invalid argument: %s",
	NotFoundError:        "Entity not found: %s",
	InternalError:        "Internal server error",
	DatabaseError:        "Database error: %s",
	NotConnectedError:    "Database connection is not established",
}

// Error represents a custom error.
type Error struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// Error returns the error message.
func (e *Error) Error() string {
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

// NewError creates a new custom error with a formatted message.
func NewError(code ErrorCode, message string) *Error {
	// Use the default message if a specific message is not provided.
	if message == "" {
		message = errorMessages[code]
	}
	return &Error{
		Code:    code,
		Message: message,
	}
}

// NewErrorWithArgs NewError creates a new custom error with a formatted message.
func NewErrorWithArgs(code ErrorCode, args ...interface{}) *Error {
	message := fmt.Sprintf(errorMessages[code], args...)

	return &Error{
		Code:    code,
		Message: message,
	}
}
