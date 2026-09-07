// Package errors (imported as apperrors) defines a business-level error type
// carrying an HTTP status and stable code, so handlers never have to guess
// how to render an error to the client. Named "errors" on disk but the Go
// package identifier is "apperrors" to avoid colliding with the stdlib
// "errors" package wherever both are imported together.
package apperrors

import (
	"errors"
	"net/http"
)

type AppError struct {
	Status  int
	Code    string
	Message string
	Err     error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return e.Message + ": " + e.Err.Error()
	}
	return e.Message
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func New(status int, code, message string, err error) *AppError {
	return &AppError{Status: status, Code: code, Message: message, Err: err}
}

func BadRequest(message string, err error) *AppError {
	return New(http.StatusBadRequest, "BAD_REQUEST", message, err)
}

func Unauthorized(message string) *AppError {
	return New(http.StatusUnauthorized, "UNAUTHORIZED", message, nil)
}

func Forbidden(message string) *AppError {
	return New(http.StatusForbidden, "FORBIDDEN", message, nil)
}

func NotFound(message string) *AppError {
	return New(http.StatusNotFound, "NOT_FOUND", message, nil)
}

func Conflict(message string) *AppError {
	return New(http.StatusConflict, "CONFLICT", message, nil)
}

func Internal(err error) *AppError {
	return New(http.StatusInternalServerError, "INTERNAL", "something went wrong", err)
}

// As extracts an *AppError from a generic error, if present.
// Named As (not the stdlib name) deliberately since callers import this
// package under the "apperrors" alias, so apperrors.As reads unambiguously.
func As(err error) (*AppError, bool) {
	var ae *AppError
	if errors.As(err, &ae) {
		return ae, true
	}
	return nil, false
}

// Is re-exports stdlib errors.Is so callers only need one import for
// business errors plus sentinel-error comparisons (e.g. user.ErrNotFound).
func Is(err, target error) bool {
	return errors.Is(err, target)
}
