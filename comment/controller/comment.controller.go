package controller

import (
	"net/http"

	"example.com/go-gin-blog-api/comment/request"
	"example.com/go-gin-blog-api/comment/service"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	service service.CommentService
}

func NewCommentController(service service.CommentService) *CommentController {
	return &CommentController{service}
}

func (cc *CommentController) CreateComment(c *gin.Context) {
	var input request.CommentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment", "details": err.Error()})
		return
	}

	username := c.GetString("username")
	result, err := cc.service.CreateComment(username, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (cc *CommentController) GetCommentsByPost(c *gin.Context) {
	postID := c.Param("id")
	comments, err := cc.service.GetCommentsByPost(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get comments"})
		return
	}
	c.JSON(http.StatusOK, comments)
}
