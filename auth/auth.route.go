package auth

import (
	"example.com/go-gin-blog-api/auth/controller"
	"example.com/go-gin-blog-api/auth/repository"
	"example.com/go-gin-blog-api/auth/service"
	"example.com/go-gin-blog-api/middleware"
	"example.com/go-gin-blog-api/utils"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine) {
	db := utils.DB

	authRepo := repository.NewAuthRepository(db)
	authService := service.NewAuthService(authRepo)
	authController := controller.NewAuthController(authService)

	auth := r.Group("/api")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.Refresh)
	}

	protected := r.Group("/api")
	protected.Use(middleware.JWTAuth())
	{
		protected.GET("/protected", authController.Protected)
	}
}
