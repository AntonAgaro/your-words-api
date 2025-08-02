package routes

import (
	"github.com/gin-gonic/gin"
	"your-words/controllers"
)

func RegisterUsersRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.GET("/", controllers.GetUsers)

	}
}
