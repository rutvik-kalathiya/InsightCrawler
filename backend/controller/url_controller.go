package controllers

import (
	"net/http"

	"insightcrawler/database"
	"insightcrawler/models"

	"github.com/gin-gonic/gin"
)

type CreateUrlInput struct {
	URL string `json:"url" binding:"required,url"`
}

func CreateUrl(c *gin.Context) {
	var input CreateUrlInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	url := models.Url{
		URL:    input.URL,
		Status: "queued",
	}

	if err := database.DB.Create(&url).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save URL"})
		return
	}

	c.JSON(http.StatusCreated, url)
}
