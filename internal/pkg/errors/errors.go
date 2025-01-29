package errors

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

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

// HTTP status codes for each error code
var errorStatusCodes = map[ErrorCode]int{
	InvalidArgumentError: http.StatusBadRequest,
	NotFoundError:        http.StatusNotFound,
	InternalError:        http.StatusInternalServerError,
	DatabaseError:        http.StatusInternalServerError,
	NotConnectedError:    http.StatusServiceUnavailable,
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

// NewError creates a new custom error.
func NewError(code ErrorCode, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

// NewErrorWithArgs creates a new custom error with formatted message.
func NewErrorWithArgs(code ErrorCode, format string, args ...interface{}) *Error {
	return &Error{
		Code:    code,
		Message: fmt.Sprintf(format, args...),
	}
}

// GetHTTPStatusCode returns the HTTP status code for an error.
func GetHTTPStatusCode(err error) int {
	var customErr *Error
	if errors.As(err, &customErr) {
		if statusCode, ok := errorStatusCodes[customErr.Code]; ok {
			return statusCode
		}
	}
	return http.StatusInternalServerError
}

// HandleHTTPError handles HTTP errors based on the custom error type.
func HandleHTTPError(c echo.Context, err error) error {
	var customErr *Error
	if errors.As(err, &customErr) {
		return c.JSON(GetHTTPStatusCode(err), customErr)
	}
	// If the error is not a custom error, return a generic internal server error.
	return c.JSON(http.StatusInternalServerError, NewError(InternalError, "Internal server error"))
}
