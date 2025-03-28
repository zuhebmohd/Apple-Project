package controllers

import (
	"net/http"
	"strconv"

	"web-app/db"
	"web-app/models"

	"github.com/gin-gonic/gin"
)

func GetClusters(c *gin.Context) {
	var clusters []models.Cluster
	db.DB.Find(&clusters)
	c.JSON(http.StatusOK, clusters)
}

func UpdateCluster(c *gin.Context) {
	// Only admin can update (enforced by admin middleware)
	id := c.Param("id")
	newCountStr := c.PostForm("server_count")
	newCount, err := strconv.Atoi(newCountStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid server count"})
		return
	}

	var cluster models.Cluster
	if err := db.DB.First(&cluster, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cluster not found"})
		return
	}

	cluster.ServerCount = newCount
	db.DB.Save(&cluster)
	c.JSON(http.StatusOK, cluster)
}
