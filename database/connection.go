package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"assignment1GO/models"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(postgres.Open("postgres://postgres:12345@localhost:5432/assignment2go"), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the db")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Item{})
}
