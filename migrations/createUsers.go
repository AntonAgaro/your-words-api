package migrations

import (
	"log"
	"your-words/database"
	"your-words/models"
)

func CreateUsersTable() {
	err := database.Db.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Error with migration CreateUserTable: ", err)
	}

	log.Println("Table Users was created!")
}
