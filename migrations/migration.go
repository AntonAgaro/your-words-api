package migrations

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"your-words/database"
)

func RunMigrations() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//handlers.AddWord()
	database.ConnectDb(os.Getenv("DATABASE_URL"))

	CreateUsersTable()
	CreateTopicTable()
	CreateWordTable()
	AddBaseTopics()

	fmt.Println("All migrations were executed!")
}
