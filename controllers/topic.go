package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"your-words/database"
	"your-words/models"
)

func GetTopics(c *gin.Context) {
	var topics []models.Topic

	if err := database.Db.Find(&topics).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error with get topics"})
		return
	}

	c.JSON(200, topics)
}
