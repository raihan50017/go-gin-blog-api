package reaction

import (
	"example.com/go-gin-blog-api/middleware"
	"example.com/go-gin-blog-api/reaction/controller"
	"example.com/go-gin-blog-api/reaction/repository"
	"example.com/go-gin-blog-api/reaction/service"
	"example.com/go-gin-blog-api/utils"
	"github.com/gin-gonic/gin"
)

func RegisterReactionRoutes(r *gin.Engine) {
	db := utils.DB

	reactionRepo := repository.NewReactionRepository(db)
	reactionService := service.NewReactionService(reactionRepo)
	reactionController := controller.NewReactionController(reactionService)

	// Public
	r.GET("/api/reactions/:id", reactionController.GetReactionsByPost)

	// Protected
	authRoutes := r.Group("/api")
	authRoutes.Use(middleware.JWTAuth())
	{
		authRoutes.POST("/reactions", reactionController.ReactToPost)
	}
}
