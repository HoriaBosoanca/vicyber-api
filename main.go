package main

import (
	"os"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"server/server"
) 
 
var router *mux.Router

func main() {
	// database conectiviy
	server.DB = server.ConnectDB()
	server.DoMigrations(server.DB)

	// route
	router = mux.NewRouter()
	router.Use(server.EnableCORS)
	server.HandleArticle(router)

	// Listen on port
	port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
	log.Printf("Server starting on port %s.", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// Handler function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
