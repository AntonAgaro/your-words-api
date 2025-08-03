package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"your-words/database"
	"your-words/models"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func GetUser(c *gin.Context) {
	userId := c.MustGet("userId").(uint)
	var user models.User
	if database.Db.First(&user, userId).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found"})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "user": ToUserResponse(user)})
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToUserResponse(user models.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
