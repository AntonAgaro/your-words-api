package routes

import (
	"github.com/gin-gonic/gin"
	"your-words/controllers"
	"your-words/middlewares"
)

func RegisterTranslateRoutes(r *gin.Engine) {
	translate := r.Group("/translate")
	{
		translate.POST("/", controllers.Translate).Use(middlewares.AuthMiddleware())
	}
}
