package services

import (
	"encoding/json"
	"log"
	"net/http"
	"your-words/database"
	"your-words/models"
)

func GetTopics(w http.ResponseWriter, r *http.Request) {
	var topics []models.Topic
	result := database.Db.Find(&topics)

	if result.Error != nil {
		log.Printf("Error with GetAll: %v", result.Error)
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(topics)
	if err != nil {
		log.Printf("Error with GetAll topics: %v", err)
		return
	}
}
