package routes

import (
	"github.com/gin-gonic/gin"
	"your-words/controllers"
	"your-words/middlewares"
)

func RegisterWordsRoutes(r *gin.Engine) {
	words := r.Group("/words").Use(middlewares.AuthMiddleware())
	{
		words.POST("/", controllers.AddWord)
		words.GET("/", controllers.GetAllWords)
	}
}
