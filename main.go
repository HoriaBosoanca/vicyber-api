package main

import (
	// logging
	"log"

	// endpoint
	"net/http"

	// router
	"os"

	"github.com/gorilla/mux"

	// rest of packages
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

	// log and serve
	port := os.Getenv("PORT") // get vercel default port
	if port == "" {
		port = "8080"
	}
	log.Println("Server starting.")
	log.Fatal(http.ListenAndServe(":"+port, router))
}

// for vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
