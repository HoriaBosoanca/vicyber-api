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
	
	r := mux.NewRouter()
	server.HandleArticle(r)

	log.Println("Server starting on http://localhost:8010")
	log.Fatal(http.ListenAndServe(":8010", r))
}
