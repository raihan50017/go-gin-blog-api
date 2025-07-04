package main

import (
	"time"

	"example.com/go-gin-blog-api/route"
	"example.com/go-gin-blog-api/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.ConnectDB()
	utils.MigrateDB()

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	route.RegisterRoutes(r)

	r.Run(":8080")
}
