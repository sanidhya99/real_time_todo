package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"fmt"
	"os"
)

var DB *gorm.DB

// ConnectDatabase connects to the PostgreSQL database using environment variables.
func ConnectDatabase() {
	// Retrieve environment variables
	host := os.Getenv("POSTGRES_HOST")          // e.g., 192.168.61.67 or localhost
	port := os.Getenv("POSTGRES_PORT")          // e.g., 5432
	user := os.Getenv("POSTGRES_USER")          // e.g., todouser
	password := os.Getenv("POSTGRES_PASSWORD")  // e.g., your password
	dbname := os.Getenv("POSTGRES_DB")          // e.g., tododb

	// Format the DSN (Data Source Name) with the retrieved environment variables
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	// Connect to the database
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Store the connection instance globally
	DB = database
}
