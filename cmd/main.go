package main

import (
	"log"
	"web-app/config"
	"web-app/db"
	"web-app/routes"
	//"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	config.LoadProperties("application.properties")

	// Connect to DB and perform migrations
	db.Connect()

	// Optionally: Seed initial data if not already present
	seedData()

	// Setup Gin router and routes
	router := gin.Default()
	routes.SetupRoutes(router)

	log.Println("Server running on port 8080...")
	router.Run(":8080")
}

func seedData() {
	// Insert default clusters if not present.
	var count int64
	db.DB.Model(&struct{ ID uint }{}).Count(&count)
	if count == 0 {
		clusters := []struct {
			Name        string
			ServerCount int
		}{
			{"Amazon", 850},
			{"Michigan", 904},
			{"Rhino", 1208},
			{"Tahoe", 150},
		}
		for _, c := range clusters {
			db.DB.Create(&struct {
				Name        string
				ServerCount int
			}{Name: c.Name, ServerCount: c.ServerCount})
		}
	}

	// Insert users based on properties file:
	var userCount int64
	db.DB.Model(&struct{ ID uint }{}).Count(&userCount)
	if userCount == 0 {
		users := []struct {
			Username string
			Password string
			Role     string
		}{
			{config.Properties["user.readonly"], config.Properties["pass.readonly"], "readonly"},
			{config.Properties["user.admin"], config.Properties["pass.admin"], "admin"},
		}
		for _, u := range users {
			db.DB.Create(&struct {
				Username string
				Password string
				Role     string
			}{Username: u.Username, Password: u.Password, Role: u.Role})
		}
	}
}
