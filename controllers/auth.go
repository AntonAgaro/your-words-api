package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
	"your-words/database"
	"your-words/models"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	if username == "" || password == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid password"})
		return
	}

	user := models.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
	}

	if err := database.Db.Create(&user).Error; err != nil {
		c.JSON(http.StatusConflict, gin.H{"status": "error", "message": "user already exists"})
		return
	}

	token := getAuthJwt(user.ID)

	if len(token) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign token"})
		return
	}

	c.SetCookie("token", token, 86400, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("User %s registered successfully", user.Username), "user": ToUserResponse(user)})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid input"})
		return
	}

	var user models.User

	if err := database.Db.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "user not found"})
		return
	}

	fmt.Println(user)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "error", "message": "incorrect username or password"})
		return
	}

	token := getAuthJwt(user.ID)

	if len(token) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign token"})
		return
	}

	c.SetCookie("token", token, 86400, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("Welcome %s!", user.Username), "user": ToUserResponse(user)})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("You successfully logout!")})

}

func getAuthJwt(userId uint) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return ""
	}

	return tokenString
}
