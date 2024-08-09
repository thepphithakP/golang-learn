package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
	RoleID   uint
	Role     Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (User) TableName() string {
	return "public.users"
}
