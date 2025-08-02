package routes

import (
	"github.com/gin-gonic/gin"
	"your-words/controllers"
)

func RegisterWordsRoutes(r *gin.Engine) {
	words := r.Group("/words")
	{
		words.POST("/", controllers.AddWord)
		words.GET("/", controllers.GetAllWords)
	}
}
