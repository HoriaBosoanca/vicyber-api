package main

import (
	// logging
	"log"

	// eviroment variables
	"os"

	// endpoint
	"encoding/json"
	"net/http"

	// GORM
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Endpoint stuff

type Response struct {
	Message string `json:"message"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := Response{
		Message: "Hello, World!",
	}
	json.NewEncoder(w).Encode(response)
}

// GORM stuff

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100"`
	Email string `gorm:"unique;size:100"`
}

func main() {
	connection := os.Getenv("VICYBERPOSTGRES")
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&User{})

	user := User{Name: "Hi Doe", Email: "hi.doe@example.com"}
	result := db.Create(&user)
	if result.Error != nil {
		log.Fatal("Error creating user:", result.Error)
	}
	log.Printf("User created with ID: %d\n", user.ID)

	http.HandleFunc("/", Handler)
	log.Println("Server starting on http://localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
