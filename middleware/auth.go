package middleware

import (
	"net/http"
	"web-app/db"
	"web-app/models"

	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing credentials"})
			return
		}

		var user models.User
		result := db.DB.Where("username = ? AND password = ?", username, password).First(&user)
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		// Save user role in context for later use
		c.Set("userRole", user.Role)
		c.Next()
	}
}
