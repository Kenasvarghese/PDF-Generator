package utils

import (
	"net/http"

	"github.com/go-playground/validator/v10"
)

// WriteApiResponse writes the status and message to the response writer
func WriteApiResponse(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(msg))
}

// validates structs
func Validate(v any) error {
	validator := validator.New()
	return validator.Struct(v)
}
