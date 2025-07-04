package model

import (
	authModel "example.com/go-gin-blog-api/auth/model"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content  string `json:"content" binding:"required"`
	PostID   uint   `json:"post_id"`
	Author   string `json:"author"`
	AuthorID uint   `json:"author_id"`

	//Post Post `json:"post" gorm:"foreignKey:PostID;references:ID"`
	User authModel.User `json:"user" gorm:"foreignKey:AuthorID;references:ID"`
}
