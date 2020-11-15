package flaw

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
