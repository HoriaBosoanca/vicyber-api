package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router *mux.Router

// InitializeRouter sets up the routes
func InitializeRouter() *mux.Router {
	Router = mux.NewRouter()
	// Add your routes
	return Router
}

// Handler function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	Router.ServeHTTP(w, r)
}
