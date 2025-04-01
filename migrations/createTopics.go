package migrations

import (
	"log"
	"your-words/database"
	"your-words/models"
)

func CreateTopicTable() {
	err := database.Db.AutoMigrate(&models.Topic{})

	if err != nil {
		log.Fatal("Error with migration CreateTopicTable")
	}

	log.Println("Table Topic was created!")
}
