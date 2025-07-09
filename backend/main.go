package main

import (
	controllers "insightcrawler/controller"
	"insightcrawler/database"
	"insightcrawler/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	database.Connect()

	// Auto-migrate your Url model
	database.DB.AutoMigrate(&models.Url{})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Add this route
	r.POST("/api/urls", controllers.CreateUrl)

	r.Run(":8080")
}
