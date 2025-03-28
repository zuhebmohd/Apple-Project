package routes

import (
	"web-app/controllers"
	"web-app/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public endpoint if needed (e.g., login)

	// Protected routes:
	api := router.Group("/api")
	api.Use(middleware.BasicAuth())
	{
		api.GET("/clusters", controllers.GetClusters)
		// Update route only accessible by admin users:
		api.PUT("/clusters/:id", middleware.AdminOnly(), controllers.UpdateCluster)
	}
}
