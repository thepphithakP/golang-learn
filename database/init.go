package database

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/thepphithakP/golang-learn/app/models"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() *gorm.DB {
    dsn := "user=devUser password=isMysecret98 dbname=userservice port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database")
    }

    // Auto Migrate
    DB.AutoMigrate(&models.User{}, &models.Role{})

    return DB
}