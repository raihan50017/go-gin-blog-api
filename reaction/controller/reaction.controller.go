package controller

import (
	"net/http"

	"example.com/go-gin-blog-api/reaction/request"
	"example.com/go-gin-blog-api/reaction/service"
	"github.com/gin-gonic/gin"
)

type ReactionController struct {
	service service.ReactionService
}

func NewReactionController(service service.ReactionService) *ReactionController {
	return &ReactionController{service}
}

func (rc *ReactionController) ReactToPost(c *gin.Context) {
	var input request.ReactionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	username := c.GetString("username")
	result, err := rc.service.ReactToPost(username, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (rc *ReactionController) GetReactionsByPost(c *gin.Context) {
	postID := c.Param("id")
	reactions, err := rc.service.GetReactionsByPost(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get reactions"})
		return
	}

	c.JSON(http.StatusOK, reactions)
}
