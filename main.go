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

func main() {
	// database conectiviy
	server.DB = server.ConnectDB()
	server.DoMigrations(server.DB)

	// route
	r := mux.NewRouter()
	server.HandleArticle(r)

	// CORS
	r.Use(server.EnableCORS)

	// log and serve
	log.Println("Server starting on http://localhost:8010")
	log.Fatal(http.ListenAndServe(":8010", r))
}
