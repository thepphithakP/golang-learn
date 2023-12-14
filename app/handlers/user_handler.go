package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thepphithakP/golang-learn/app/services"
	"github.com/thepphithakP/golang-learn/database/repositories"
	"log"
	"strconv"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userRepository *repositories.UserRepository) *UserHandler {
	userService := services.NewUserService(userRepository)
	return &UserHandler{
		UserService: userService,
	}
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	userIDParam := c.Params("id")
	log.Printf("userIDParam : " + userIDParam)
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		// Log error
		log.Printf("Failed to parse user ID: %v", err)
		// Return HTTP 400 Bad Request response
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID",
		})
	}
	log.Printf("userID: %d", userID)

	user, err := h.UserService.GetUserByID(uint(userID))
	if err != nil {
		// Log error
		log.Printf("Failed to retrieve user: %v", err)
		// Return HTTP 500 Internal Server Error response
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve user",
		})
		return nil // ไม่ return error นี้
	}

	if user == nil {
		// Handle the case where user is nil
		log.Println("User is nil")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "User is nil",
		})
	} else {
		return c.JSON(user)
	}

}
