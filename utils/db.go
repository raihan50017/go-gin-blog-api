package utils

import (
	"fmt"
	"log"
	"os"

	authModel "example.com/go-gin-blog-api/auth/model"
	commentModel "example.com/go-gin-blog-api/comment/model"
	postModel "example.com/go-gin-blog-api/post/model"
	reactionModel "example.com/go-gin-blog-api/reaction/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	DB = db
	fmt.Println("Connected to PostgreSQL")
}

func MigrateDB() {
	DB.AutoMigrate(&authModel.User{}, &postModel.Post{}, &commentModel.Comment{}, &reactionModel.Reaction{}, &authModel.RefreshToken{})
}
