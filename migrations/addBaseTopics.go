package migrations

import (
	"log"
	"your-words/database"
	"your-words/models"
)

func AddBaseTopics() {
	var count int64
	database.Db.Model(&models.Topic{}).Count(&count)

	if count > 0 {
		log.Fatal("Base topics were already added!")
		return
	}

	topics := []*models.Topic{
		{Name: "House", Description: "Words about home and furniture"},
		{Name: "Family", Description: "Words about family and relationships"},
		{Name: "Food", Description: "Words about meals, fruits, and vegetables"},
		{Name: "Health", Description: "Words about body, fitness, and medicine"},
		{Name: "Clothing", Description: "Words about clothes, shoes, and accessories"},
		{Name: "Transport", Description: "Words about cars, buses, and planes"},
		{Name: "Nature", Description: "Words about weather, animals, and plants"},
		{Name: "Work & Study", Description: "Words about jobs, school, and office"},
		{Name: "Entertainment", Description: "Words about music, movies, and books"},
		{Name: "Finance", Description: "Words about shopping, money, and banking"},
	}

	result := database.Db.Create(topics)

	if result.Error != nil {
		log.Fatal("Error with insert base topics")
	}

	log.Printf("%v were inserted in table topic!", result.RowsAffected)
}
