package main

import (
	// logging
	"log"

	// endpoint
	"net/http"

	// router
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

	log.Println("Server starting.")
	log.Fatal(http.ListenAndServe(":8010", router))
}

// Handler function for Vercel
func Handler(w http.ResponseWriter, r *http.Request) {
	router.ServeHTTP(w, r)
}
