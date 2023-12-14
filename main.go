package main

import (
	"database/sql"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/thepphithakP/golang-learn/database"
	"github.com/thepphithakP/golang-learn/routes"
)

func main() {
	app := fiber.New()
	// Initialize database
	DB := database.InitDB()
	// Setup routes
	routes.SetupRoutes(app, DB)
	db, err := database.DB.DB()

	if err != nil {
		fmt.Println("Error connecting database")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	err = app.Listen(":3000")
	if err != nil {
		return
	}
}
