package db

import (
	"log"
	"web-app/config"
	"web-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	dsn := config.Properties["db.url"]
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	// Auto-migrate your models
	err = DB.AutoMigrate(&models.Cluster{}, &models.User{})
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}
}
