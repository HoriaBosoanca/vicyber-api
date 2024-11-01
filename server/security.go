package server

import (
	"net/http"
	"os"
	"log"
)

func CheckApiKey(w http.ResponseWriter, r *http.Request) bool {
	apiKey := r.Header.Get("Authorization")
	if apiKey != os.Getenv("VICYBERAPIKEY") {
		http.Error(w, "Unauthorized: Wrong API key!", http.StatusUnauthorized)
		logApiKeyAttempt(apiKey)
		return false
	}
	logApiKeyAttempt("Correct api key!")
	return true
}

func logApiKeyAttempt(s string) {
	logFile, err := os.OpenFile("apikeylog.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Error opening logfile: %v", err)
		return
	}
	defer logFile.Close()

	logger := log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println(s)
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    w.WriteHeader(http.StatusOK)
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
