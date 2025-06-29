package controllers

import (
	"net/http"
	"time"

	"example.com/go-gin-blog-api/models"
	"example.com/go-gin-blog-api/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

const (
	accessTokenDuration  = 15 * time.Minute
	refreshTokenDuration = 24 * time.Hour
)

type AuthController struct {
	DB *gorm.DB
}

func (ac *AuthController) Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var existing models.User
	if err := ac.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already registered"})
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Password hashing failed"})
		return
	}
	input.Password = string(hashed)

	if err := ac.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User creation failed"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (ac *AuthController) Login(c *gin.Context) {
	var input models.User
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ac.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	accessToken, _ := utils.GenerateToken(user.Username, accessTokenDuration)
	refreshToken, _ := utils.GenerateToken(user.Username, refreshTokenDuration)

	ac.DB.Create(&models.RefreshToken{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(refreshTokenDuration).Unix(),
	})

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
		"user": gin.H{
			"username": user.Username,
			"email":    user.Email,
			"id":       user.ID,
		},
	})
}

func (ac *AuthController) Refresh(c *gin.Context) {
	var body map[string]string
	if err := c.ShouldBindJSON(&body); err != nil || body["refresh_token"] == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token required"})
		return
	}

	tokenStr := body["refresh_token"]
	claims, err := utils.ParseToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	var stored models.RefreshToken
	if err := ac.DB.Where("token = ?", tokenStr).First(&stored).Error; err != nil || time.Now().Unix() > stored.ExpiresAt {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Refresh token invalid or expired"})
		return
	}

	ac.DB.Delete(&stored)

	accessToken, _ := utils.GenerateToken(claims.Username, accessTokenDuration)
	newRefreshToken, _ := utils.GenerateToken(claims.Username, refreshTokenDuration)

	var user models.User
	ac.DB.Where("username = ?", claims.Username).First(&user)

	ac.DB.Create(&models.RefreshToken{
		Token:     newRefreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(refreshTokenDuration).Unix(),
	})

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": newRefreshToken,
	})
}

func (ac *AuthController) Protected(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{
		"message":  "Welcome to the protected route!",
		"username": username,
	})
}
