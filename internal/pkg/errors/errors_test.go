package errors

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	err := NewError(InvalidArgumentError, "test message")
	assert.Equal(t, "[InvalidArgument] test message", err.Error())
}

func TestNewError(t *testing.T) {
	err := NewError(InvalidArgumentError, "test message")
	assert.Equal(t, ErrorCode("InvalidArgument"), err.Code)
	assert.Equal(t, "test message", err.Message)
}

func TestNewErrorWithArgs(t *testing.T) {
	err := NewErrorWithArgs(DatabaseError, "test %s", "message")
	assert.Equal(t, ErrorCode("DatabaseError"), err.Code)
	assert.Equal(t, "test message", err.Message)
}

func TestHandleHTTPError_CustomError(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := HandleHTTPError(c, NewErrorWithArgs(NotFoundError, "test %s", "message"))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestHandleHTTPError_Default(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	err := HandleHTTPError(c, fmt.Errorf("some error"))

	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetHTTPStatusCode(t *testing.T) {
	// Test for a known error code.
	code := GetHTTPStatusCode(NewError(InvalidArgumentError, ""))
	assert.Equal(t, http.StatusBadRequest, code)

	// Test for an unknown error code.
	code = GetHTTPStatusCode(NewError(ErrorCode("UnknownErrorCode"), ""))
	assert.Equal(t, http.StatusInternalServerError, code)

	// Test for a non-custom error.
	code = GetHTTPStatusCode(fmt.Errorf("some error"))
	assert.Equal(t, http.StatusInternalServerError, code)
}
