package api

import (
	"encoding/json"
	_ "encoding/json"
	"net/http"
	_ "net/http"
)

// Coin balance param
type CoinBalanceParams struct {
	Username string
}

// Coin Balance Response
type CoinBalanceResponse struct {
	// Success Code, usually  200
	Code int
	// Account balance
	Balance int64
}

// Error response
type Error struct {
	// Error code
	Code int
	// Error message
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}

	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An expected Error Occurred.", http.StatusInternalServerError)
	}
)
