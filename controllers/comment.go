package controllers

import (
	"net/http"

	"example.com/go-gin-blog-api/dtos"
	"example.com/go-gin-blog-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	DB *gorm.DB
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var input dtos.CommentInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment", "details": err.Error()})
		return
	}

	username := c.GetString("username")

	var user models.User
	if err := cc.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	comment := models.Comment{
		Content:  input.Content,
		PostID:   input.PostID,
		Author:   user.Username,
		AuthorID: user.ID,
	}

	if err := cc.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create comment"})
		return
	}

	if err := cc.DB.Preload("User").First(&comment, comment.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load comment user"})
		return
	}

	commentResponse := dtos.ToCommentResponse(comment)

	c.JSON(http.StatusCreated, commentResponse)
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
