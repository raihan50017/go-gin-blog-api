package controllers

import (
	"net/http"

	"example.com/go-gin-blog-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BlogController struct {
	DB *gorm.DB
}

func (bc *BlogController) CreatePost(c *gin.Context) {

	var input models.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	username := c.GetString("username")
	var user models.User
	if err := bc.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	input.Author = username
	input.AuthorID = user.ID

	if err := bc.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func (bc *BlogController) GetPosts(c *gin.Context) {
	var posts []models.Post
	bc.DB.Order("created_at desc").Find(&posts)
	c.JSON(http.StatusOK, posts)
}

func (bc *BlogController) GetPostByID(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	if err := bc.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (bc *BlogController) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	username := c.GetString("username")

	var post models.Post
	if err := bc.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.Author != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this post"})
		return
	}

	var updated models.Post
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	post.Title = updated.Title
	post.Content = updated.Content
	bc.DB.Save(&post)

	c.JSON(http.StatusOK, post)
}

func (bc *BlogController) DeletePost(c *gin.Context) {
	id := c.Param("id")
	username := c.GetString("username")

	var post models.Post
	if err := bc.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	if post.Author != username {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not the author of this post"})
		return
	}

	bc.DB.Delete(&post)
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}
