package routes

import (
	"github.com/gin-gonic/gin"
	"your-words/controllers"
	"your-words/middlewares"
)

func RegisterUsersRoutes(r *gin.Engine) {
	users := r.Group("/users").Use(middlewares.AuthMiddleware())
	{
		users.GET("/", controllers.GetUsers)
		users.GET("/me", controllers.GetUser)
	}
}
