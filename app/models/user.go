package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username string    `gorm:"unique"`
	Email    string    `gorm:"unique"`
	Password string
	RoleID   uint
	Role     Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (User) TableName() string {
	return "demo.users"
}
