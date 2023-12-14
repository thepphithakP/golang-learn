// routes/routes.go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thepphithakP/golang-learn/app/handlers"
	"github.com/thepphithakP/golang-learn/database/repositories"
	"gorm.io/gorm"
)

// type AppHandlers struct {
// 	UserHandler *handlers.UserHandler
// }

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	userRepository := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepository)
	app.Get("/user/:id", userHandler.GetUserByID)
}
