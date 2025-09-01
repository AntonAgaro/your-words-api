package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	RegisterWordsRoutes(r)
	RegisterTopicsRoutes(r)
	RegisterAuthRoutes(r)
	RegisterUsersRoutes(r)
	RegisterTranslateRoutes(r)
}
