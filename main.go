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
	"server/handler"
)

func main() {
	// database conectiviy
	server.DB = server.ConnectDB()
	server.DoMigrations(server.DB)

	// route
	handler.Router = mux.NewRouter()
	handler.Router.Use(server.EnableCORS)
	server.HandleArticle(handler.Router)

	// log and serve
	port := os.Getenv("PORT") // get vercel default port
	if port == "" {
		port = "8080"
	}
	log.Println("Server starting.")
	log.Fatal(http.ListenAndServe(":"+port, handler.Router))
}
