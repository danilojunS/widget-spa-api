package handlers

import (
	"net/http"
)

func buildErrorResolver(message string, status int) func(w http.ResponseWriter, customMessage string) {
	return func(w http.ResponseWriter, customMessage string) {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(status)

		if customMessage != "" {
			message = customMessage
		}

		_, err := w.Write([]byte(message))
		if err != nil {
			panic(err)
		}
	}
}

// InternalError handles unexpected server errors
var InternalError = buildErrorResolver("Internal server error", http.StatusInternalServerError)

// ValidationError handles parameters validation erros
var ValidationError = buildErrorResolver("Params validation error", http.StatusBadRequest)

// NotFoundError handles not found errors
var NotFoundError = buildErrorResolver("Not found", http.StatusNotFound)

// UnauthorizedError handles not found errors
var UnauthorizedError = buildErrorResolver("Unauthorized", http.StatusUnauthorized)
