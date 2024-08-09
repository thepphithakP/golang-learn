package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thepphithakP/golang-learn/app/models"
	"github.com/thepphithakP/golang-learn/app/services"
	"github.com/thepphithakP/golang-learn/database/repositories"
	"log"
)

type RoleHandler struct {
	RoleService *services.RoleService
}

func NewRoleHandler(roleRepo *repositories.RoleRepository) *RoleHandler {
	roleService := services.NewRoleService(roleRepo)
	return &RoleHandler{
		RoleService: roleService,
	}
}

type RoleRequest struct {
	Name string `json:"name"`
}

func (h *RoleHandler) CreateRole(c *fiber.Ctx) error {
	payload := new(RoleRequest)
	if err := c.BodyParser(payload); err != nil {
		// Log error
		log.Printf("Failed to parse request body: %v", err)
		// Return HTTP 400 Bad Request response
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	role := &models.Role{Name: payload.Name}

	_, err := h.RoleService.Create(role)

	if err != nil {
		// Log error
		log.Printf("Failed to create role: %v", err)
		// Return HTTP 500 Internal Server Error response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create role",
		})
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *RoleHandler) GetRoles(c *fiber.Ctx) error {
	roles, err := h.RoleService.GetRoles()
	if err != nil {
		// Log error
		log.Printf("Failed to retrieve roles: %v", err)
		// Return HTTP 500 Internal Server Error response
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve roles",
		})
	}

	return c.JSON(roles)

}
