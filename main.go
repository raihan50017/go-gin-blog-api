package main

import (
	"example.com/go-gin-blog-api/controllers"
	"example.com/go-gin-blog-api/middleware"
	"example.com/go-gin-blog-api/models"
	"example.com/go-gin-blog-api/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.ConnectDB()
	utils.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Reaction{}, &models.RefreshToken{})

	auth := &controllers.AuthController{DB: utils.DB}
	blog := &controllers.BlogController{DB: utils.DB}
	comment := &controllers.CommentController{DB: utils.DB}
	reaction := &controllers.ReactionController{DB: utils.DB}

	r := gin.Default()

	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)
	r.POST("/refresh", auth.Refresh)

	r.GET("/api/posts", blog.GetPosts)
	r.GET("/api/posts/:id", blog.GetPostByID)
	r.GET("/api/posts/:id/comments", comment.GetCommentsByPost)
	r.GET("/api/posts/:id/reactions", reaction.GetReactionsByPost)

	// Protected actions
	authRoutes := r.Group("/api")
	authRoutes.Use(middleware.JWTAuth())
	{
		authRoutes.GET("/protected", auth.Protected)
		authRoutes.POST("/posts", blog.CreatePost)
		authRoutes.PUT("/posts/:id", blog.UpdatePost)
		authRoutes.DELETE("/posts/:id", blog.DeletePost)

		authRoutes.POST("/comments", comment.CreateComment)
		authRoutes.POST("/reactions", reaction.ReactToPost)
	}

	r.Run(":8080")
}
