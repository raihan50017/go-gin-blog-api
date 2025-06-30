package controllers

import (
	"net/http"

	"example.com/go-gin-blog-api/dtos"
	"example.com/go-gin-blog-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ReactionController struct {
	DB *gorm.DB
}

func (rc *ReactionController) ReactToPost(c *gin.Context) {

	var input dtos.ReactionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	username := c.GetString("username")

	var user models.User
	if err := rc.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	var existing models.Reaction
	err := rc.DB.Where("post_id = ? AND author_id = ?", input.PostID, user.ID).First(&existing).Error

	if err == nil {
		existing.Type = input.Type
		rc.DB.Save(&existing)
		c.JSON(http.StatusOK, existing)
		return
	}

	reaction := models.Reaction{
		Type:     input.Type,
		PostID:   input.PostID,
		Author:   user.Username,
		AuthorID: user.ID,
	}

	if err := rc.DB.Create(&reaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to react"})
		return
	}

	c.JSON(http.StatusCreated, reaction)
}

func (rc *ReactionController) GetReactionsByPost(c *gin.Context) {
	postID := c.Param("id")
	var reactions []models.Reaction

	if err := rc.DB.Where("post_id = ?", postID).Find(&reactions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reactions"})
		return
	}

	c.JSON(http.StatusOK, reactions)
}
