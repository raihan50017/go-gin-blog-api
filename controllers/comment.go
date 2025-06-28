package controllers

import (
	"net/http"

	"example.com/go-gin-blog-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var input models.Comment
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment"})
		return
	}

	username := c.GetString("username")
	var user models.User
	if err := cc.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	input.Author = username
	input.AuthorID = user.ID

	if err := cc.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (cc *CommentController) GetCommentsByPost(c *gin.Context) {
	postID := c.Param("id")
	var comments []models.Comment

	if err := cc.DB.Where("post_id = ?", postID).Order("created_at").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}

	c.JSON(http.StatusOK, comments)
}
