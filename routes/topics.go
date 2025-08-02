package routes

import (
	"github.com/gin-gonic/gin"
	"your-words/controllers"
)

func RegisterTopicsRoutes(r *gin.Engine) {
	topics := r.Group("/topics")
	{
		topics.GET("/", controllers.GetTopics)
	}
}
