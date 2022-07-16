package errors

import (
	"context"
	"encoding/json"
	"net/http"
	"runtime"
)

func (e GeneralError) Error() string {
	return "General Error"
}

type GeneralError struct {
	Code             int         `json:"code"`
	Message          string      `json:"message"`
	File             string      `json:"file"`
	Trace            string      `json:"trace"`
	Line             int         `json:"line"`
	Type             string      `json:"type"`
}

// New domain error
func NewDomainError(message string, code int, trace string) GeneralError {
	path, line := filePath()

	return GeneralError{
		Code:             code,
		Message:          message,
		File:             path,
		Trace:            trace,
		Line:             line,
		Type:             "DOMAIN",
	}
}

// new application error
func NewApplicationError(message string, code int, trace string) GeneralError {
	path, line := filePath()

	return GeneralError{
		Code:             code,
		Message:          message,
		File:             path,
		Trace:            trace,
		Line:             line,
		Type:             "APPLICATION",
	}
}

// new validation error
func NewValidationError(message string, code int, trace string) GeneralError {
	path, line := filePath()

	return GeneralError{
		Code:             code,
		Message:          message,
		File:             path,
		Trace:            trace,
		Line:             line,
		Type:             "VALIDATION",
	}

}


// Error Encoder
func EncodeGeneralErrorResponse(_ context.Context, err error, w http.ResponseWriter) {

	res := err.(GeneralError)
	w.Header().Set("Content-Type", "application/json")
	httpStatus := http.StatusInternalServerError
	if res.Type == "VALIDATION" {
		httpStatus = http.StatusUnprocessableEntity
	} else if res.Type == "DOMAIN" {
		httpStatus = http.StatusBadRequest
	}
	w.WriteHeader(httpStatus)

	json.NewEncoder(w).Encode(res)
}

func filePath() (string, int) {
	_, file, line, ok := runtime.Caller(2)
	if ok {
		return file, line
	}

	return "", 0
}