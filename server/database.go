package server

import (
	// logging
	"log"

	// eviroment variables
	"os"

	// GORM
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	connection := os.Getenv("POSTGRES_URL")
	db, err := gorm.Open(postgres.Open(connection), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func DoMigrations(db *gorm.DB) {
	db.AutoMigrate(&Article{})
}
