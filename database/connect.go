package database

import (
	"Cargo_Dash/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(postgres.Open("host=localhost user=cargo password=cargo dbname=cargo port=5432"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, models.Cargo{})

}
