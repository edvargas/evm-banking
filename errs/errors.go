package errs

import "net/http"

type AppError struct {
	Code    int    `json:",omitempty"`
	Message string `json:"message"`
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		http.StatusNotFound,
		message,
	}
}

func NewNoContent() *AppError {
	return &AppError{
		http.StatusNoContent,
		"",
	}
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		http.StatusInternalServerError,
		message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		http.StatusUnprocessableEntity,
		message,
	}
}
