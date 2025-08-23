package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"your-words/database"
	"your-words/models"
)

func AddWord(c *gin.Context) {
	userId := c.MustGet("userId").(uint)

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
	if database.Db.Where("user_id = ? AND text = ? AND topic_id = ?", userId, word.Text, word.TopicID).First(&wordExist).Error == nil {
		c.JSON(http.StatusConflict, gin.H{"status": "error", "message": fmt.Sprintf("Word %s already exist", word.Text)})
		return
	}

	word.UserID = userId

	if err := database.Db.Create(&word).Error; err != nil {
		log.Printf("Error with create word %v", word.Text)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error with create word"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": fmt.Sprintf("Word %s created", word.Text)})
	fmt.Printf("Word %s created", word.Text)
}

func GetAllWords(c *gin.Context) {
	userId := c.MustGet("userId").(uint)

	var words *[]models.Word
	if err := database.Db.Where("user_id = ?", userId).Find(&words).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Error with get words"})
		return
	}

	c.JSON(http.StatusOK, words)
}

func GetRandomWords(c *gin.Context) {
	userId := c.MustGet("userId").(uint)
	limitParam := c.Query("limit")
	limit := 5
	if limitParam != "" {
		if v, err := strconv.Atoi(limitParam); err == nil {
			if v > 1 && v <= 50 {
				limit = v
			} else {
				c.JSON(http.StatusBadRequest, gin.H{
					"status":  "error",
					"message": "limit must be between 1 and 50",
				})
				return
			}
		}
	}
	var words *[]models.Word
	if err := database.Db.Where("user_id = ?", userId).Order("RANDOM()").Limit(limit).Find(&words).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "Error with finding random words!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Random words were successfully found!", "words": words})
}
