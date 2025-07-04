package controller

import (
	"net/http"

	"example.com/go-gin-blog-api/post/model"
	"example.com/go-gin-blog-api/post/response"
	"example.com/go-gin-blog-api/post/service"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	service service.PostService
}

func NewPostController(service service.PostService) *PostController {
	return &PostController{service}
}

func (bc *PostController) CreatePost(c *gin.Context) {
	var input model.Post
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	username := c.GetString("username")
	post, err := bc.service.CreatePost(input, username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
}

func (bc *PostController) GetPosts(c *gin.Context) {
	posts, err := bc.service.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var responses []response.PostResponse
	for _, post := range posts {
		responses = append(responses, response.ToPostResponse(post))
	}

	c.JSON(http.StatusOK, responses)
}

func (bc *PostController) GetPostByID(c *gin.Context) {
	id := c.Param("id")
	post, err := bc.service.GetPostByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	c.JSON(http.StatusOK, post)
}

func (bc *PostController) UpdatePost(c *gin.Context) {
	id := c.Param("id")
	username := c.GetString("username")

	var updated model.Post
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	post, err := bc.service.UpdatePost(id, username, updated)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "unauthorized" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, post)
}

func (bc *PostController) DeletePost(c *gin.Context) {
	id := c.Param("id")
	username := c.GetString("username")

	err := bc.service.DeletePost(id, username)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "unauthorized" {
			status = http.StatusForbidden
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Post deleted"})
}

func (bc *PostController) GetMyPosts(c *gin.Context) {
	username := c.GetString("username")
	posts, err := bc.service.GetUserPosts(username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, posts)
}
