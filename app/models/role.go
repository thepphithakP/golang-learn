package models

import (
    "gorm.io/gorm"
)

type Role struct {
    gorm.Model
    Name string `gorm:"unique"`
}

func (Role) TableName() string {
    return "demo.roles"
}