package errs

import (
	"fmt"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:",omitempty"`
	Message    string `json:"message"`
	ErrorId    int    `json:"errorId"`
	Code       string
	UserID     int
	Errors     []string
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

type Errors struct {
	Code    string
	Message string
}

func (e AppError) Error() string {
	return e.Message
}
func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewUnexpectedError() *AppError {
	return &AppError{
		Message:    "Oops! the Server encountered a temporary error and could not complete your request",
		Code:       "500",
		StatusCode: http.StatusInternalServerError,
	}
}

func NewUnexpectedErrorWithMsg(msg string) *AppError {
	return &AppError{
		Message: fmt.Sprintf("Unexpected error: %s", msg),
	}
}

func ValidateResponse(Errors []string, statusCode int) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Errors:     Errors,
	}
}
