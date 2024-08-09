package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thepphithakP/golang-learn/app/handlers"
	"github.com/thepphithakP/golang-learn/database/repositories"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	// users
	userRepository := repositories.NewUserRepository(db)
	userHandler := handlers.NewUserHandler(userRepository)
	app.Get("/user/:id", userHandler.GetUserByID)
	app.Post("/user", userHandler.CreateUser)

	// roles
	roleRepository := repositories.NewRoleRepository(db)
	roleHandler := handlers.NewRoleHandler(roleRepository)
	app.Post("/role", roleHandler.CreateRole)
	app.Get("/roles", roleHandler.GetRoles)
}
