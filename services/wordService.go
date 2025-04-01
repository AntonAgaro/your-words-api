package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"your-words/database"
	"your-words/models"
	"your-words/utils"
)

func AddWord(w http.ResponseWriter, r *http.Request) {
	var word models.Word
	err := json.NewDecoder(r.Body).Decode(&word)

	if err != nil {
		utils.WriteJson(w, map[string]string{"status": "error", "message": "Error with decode word"})
		return
	}

	if word.Translation == "" || word.Text == "" {
		utils.WriteJson(w, map[string]string{"status": "error", "message": "Text and translation are required"})
		return
	}

	//Check if word topic exist
	var topic models.Topic
	if database.Db.First(&topic, word.TopicID).Error != nil {
		utils.WriteJson(w, map[string]string{"status": "error", "message": "Topic not found"})
		return
	}

	//Check if word already exist
	var wordExist models.Word
	if database.Db.Where("text = ? AND topic_id = ?", word.Text, word.TopicID).First(&wordExist).Error == nil {
		utils.WriteJson(w, map[string]string{"status": "error", "message": fmt.Sprintf("Word %s already exist", word.Text)})
		return
	}

	result := database.Db.Create(&word)
	if result.Error != nil {
		log.Printf("Error with create word %v", word.Text)
		return
	}

	utils.WriteJson(w, map[string]string{"status": "success", "message": fmt.Sprintf("Word %s created", word.Text)})

}

func GetAllWords(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetWords")
	var words *[]models.Word
	result := database.Db.Find(&words)
	if result.Error != nil {
	}
	err := json.NewEncoder(w).Encode(words)
	if err != nil {
		return
	}
}
