package server

import (
	"net/http"
	"os"
)

func CheckApiKey(w http.ResponseWriter, r *http.Request) bool {
	apiKey := r.Header.Get("Authorization")
	if apiKey != os.Getenv("VICYBERAPIKEY") {
		http.Error(w, "Unauthorized: Wrong API key!", http.StatusUnauthorized)
		return false
	}
	return true
}
