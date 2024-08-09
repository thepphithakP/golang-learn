package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/thepphithakP/golang-learn/database"
	"github.com/thepphithakP/golang-learn/routes"
	"log"
	"os"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	app := fiber.New()

	// Initialize database
	initDatabase := database.InitDB()

	// Setup routes
	routes.SetupRoutes(app, initDatabase)

	// Get the underlying sql.DB object
	db, err := database.DB.DB()
	if err != nil {
		log.Fatalf("Error getting database object: %v", err)
	}
	defer closeDatabase(db)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func closeDatabase(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
}
