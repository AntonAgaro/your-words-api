package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"your-words/database"
	"your-words/models"
)

func AddWord(c *gin.Context) {
	var word models.Word
	if err := c.ShouldBindJSON(&word); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error with decode word"})
		return
	}

	if word.Translation == "" || word.Text == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Text and translation are required"})
		return
	}

	//Check if word topic exist
	var topic models.Topic
	if database.Db.First(&topic, word.TopicID).Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Topic not found"})
		return
	}

	//Check if word already exist
	var wordExist models.Word
	if database.Db.Where("text = ? AND topic_id = ?", word.Text, word.TopicID).First(&wordExist).Error == nil {
		c.JSON(http.StatusConflict, gin.H{"status": "error", "message": fmt.Sprintf("Word %s already exist", word.Text)})
		return
	}

	if err := database.Db.Create(&word).Error; err != nil {
		log.Printf("Error with create word %v", word.Text)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error with create word"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": fmt.Sprintf("Word %s created", word.Text)})
	fmt.Printf("Word %s created", word.Text)
}

func GetAllWords(c *gin.Context) {
	var words *[]models.Word
	if err := database.Db.Find(&words).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error with get words"})
		return
	}

	c.JSON(http.StatusOK, words)
}
