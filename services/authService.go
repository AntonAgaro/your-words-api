package services

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
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid input"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid password"})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: string(hashedPassword),
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

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("User %s registered successfully", user.Username)})
}

func Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "invalid input"})
		return
	}

	var user models.User

	if err := database.Db.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "user not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "incorrect username or password"})
		return
	}

	token := getAuthJwt(user.ID)

	if len(token) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to sign token"})
		return
	}

	c.SetCookie("token", token, 86400, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": fmt.Sprintf("Welcome %s!", user.Username)})
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
