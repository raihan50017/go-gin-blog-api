package post

import (
	"example.com/go-gin-blog-api/middleware"
	"example.com/go-gin-blog-api/post/controller"
	"example.com/go-gin-blog-api/post/repository"
	"example.com/go-gin-blog-api/post/service"
	"example.com/go-gin-blog-api/utils"
	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine) {
	db := utils.DB

	postRepo := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepo)
	postController := controller.NewPostController(postService)

	// Public routes
	r.GET("/api/posts", postController.GetPosts)
	r.GET("/api/posts/:id", postController.GetPostByID)

	// Protected routes
	authRoutes := r.Group("/api")
	authRoutes.Use(middleware.JWTAuth())
	{
		authRoutes.POST("/posts", postController.CreatePost)
		authRoutes.PUT("/posts/:id", postController.UpdatePost)
		authRoutes.DELETE("/posts/:id", postController.DeletePost)
		authRoutes.GET("/myposts", postController.GetMyPosts)
	}
}
