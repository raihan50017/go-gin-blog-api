package comment

import (
	"example.com/go-gin-blog-api/comment/controller"
	"example.com/go-gin-blog-api/comment/repository"
	"example.com/go-gin-blog-api/comment/service"
	"example.com/go-gin-blog-api/middleware"
	"example.com/go-gin-blog-api/utils"
	"github.com/gin-gonic/gin"
)

func RegisterCommentRoutes(r *gin.Engine) {
	db := utils.DB

	commentRepo := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepo)
	commentController := controller.NewCommentController(commentService)

	r.GET("/api/comments/:id", commentController.GetCommentsByPost)

	authRoutes := r.Group("/api")
	authRoutes.Use(middleware.JWTAuth())
	{
		authRoutes.POST("/comments", commentController.CreateComment)
	}
}
