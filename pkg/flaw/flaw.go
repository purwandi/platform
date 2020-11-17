package flaw

import "net/http"

// AppError ....
type AppError struct {
	StatusCode int
	Method     string
	Message    string
}

func (e *AppError) Error() string {
	return e.Message
}

// Error ...
func Error(code int, message string) *AppError {
	return &AppError{
		StatusCode: code,
		Message:    message,
	}
}

// InternalError ...
func InternalError(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Message:    message,
	}
}

// BadRequest ...
func BadRequest(message string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Message:    message,
	}
}
