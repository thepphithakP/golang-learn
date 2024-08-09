package services

import (
	"github.com/thepphithakP/golang-learn/app/models"
	"github.com/thepphithakP/golang-learn/database/repositories"
)

type RoleService struct {
	roleRepository *repositories.RoleRepository
}

func NewRoleService(roleRepository *repositories.RoleRepository) *RoleService {
	return &RoleService{
		roleRepository: roleRepository,
	}
}

func (s *RoleService) GetRoles() ([]*models.Role, error) {
	return s.roleRepository.GetRoles()
}

func (s *RoleService) Create(role *models.Role) (*models.Role, error) {
	return s.roleRepository.Create(role)
}
