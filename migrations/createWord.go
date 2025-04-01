package migrations

import (
	"log"
	"your-words/database"
	"your-words/models"
)

func CreateWordTable() {
	err := database.Db.AutoMigrate(&models.Word{})

	if err != nil {
		log.Fatal("Error with migration CreateWordTable")
	}

	log.Println("Table Word was created!")
}
