package repositories

import (
	"github.com/thepphithakP/golang-learn/app/models"
	"gorm.io/gorm"
	"log"
)

type RoleRepository struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (r *RoleRepository) GetRoles() ([]*models.Role, error) {
	var roles []*models.Role
	if err := r.db.Find(&roles).Error; err != nil {
		// Log error
		log.Printf("Error retrieving roles: %v", err)
		return nil, err
	}
	return roles, nil
}

func (r *RoleRepository) Create(role *models.Role) (*models.Role, error) {
	if err := r.db.Create(role).Error; err != nil {
		// Log error
		log.Printf("Error creating role: %v", err)
		return nil, err
	}
	return role, nil
}
