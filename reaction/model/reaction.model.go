package model

import (
	authModel "example.com/go-gin-blog-api/auth/model"
	"gorm.io/gorm"
)

type Reaction struct {
	gorm.Model
	PostID   uint   `json:"post_id"`
	AuthorID uint   `json:"author_id"`
	Author   string `json:"author"`
	Type     string `json:"type"`

	//Post model.Post `json:"post" gorm:"foreignKey:PostID;references:ID"`
	User authModel.User `json:"user" gorm:"foreignKey:AuthorID;references:ID"`
}
