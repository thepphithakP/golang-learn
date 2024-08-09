package services

import (
	"github.com/thepphithakP/golang-learn/app/models"
	"github.com/thepphithakP/golang-learn/database/repositories"
)

type UserService struct {
	userRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUserByID(userID uint) (*models.User, error) {
	return s.userRepository.GetUserByID(userID)
}

func (s *UserService) Create(user *models.User) (*models.User, error) {
	createdUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}
