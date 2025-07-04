package controller

import (
	"net/http"

	"example.com/go-gin-blog-api/auth/model"
	"example.com/go-gin-blog-api/auth/service"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	service service.AuthService
}

func NewAuthController(s service.AuthService) *AuthController {
	return &AuthController{service: s}
}

func (ac *AuthController) Register(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := ac.service.Register(&input); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Registration successful"})
}

func (ac *AuthController) Login(c *gin.Context) {
	var input model.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	access, refresh, user, err := ac.service.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
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

	newAccess, newRefresh, err := ac.service.Refresh(body["refresh_token"])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  newAccess,
		"refresh_token": newRefresh,
	})
}

func (ac *AuthController) Protected(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(http.StatusOK, gin.H{
		"message":  "Welcome to the protected route!",
		"username": username,
	})
}
