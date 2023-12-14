package repositories

import (
	"github.com/thepphithakP/golang-learn/app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		// Log error
		log.Printf("Error in GetUserByID: %v", err)
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	// Hash the password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		// Log error
		log.Printf("Error hashing password: %v", err)
		return err
	}

	// Set the hashed password to the user
	user.Password = string(hashedPassword)

	// Create the user in the database
	if err := r.db.Create(user).Error; err != nil {
		// Log error
		log.Printf("Error creating user: %v", err)
		return err
	}

	return nil
}
