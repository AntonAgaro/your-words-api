package main

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"your-words/database"
	"your-words/services"
)

func main() {
	if os.Getenv("APP_ENV") != "prod" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	database.ConnectDb(os.Getenv("DATABASE_URL"))

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{fmt.Sprintf("%s:%s", os.Getenv("FRONT_HOST"), os.Getenv("FRONT_PORT"))},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/words", services.AddWord)
	router.GET("/words", services.GetAllWords)
	router.GET("/topics", services.GetTopics)

	if err := router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT"))); err != nil {
		log.Fatal(err)
	}

}
