package model

import (
	"example.com/go-gin-blog-api/auth/model"
	commentModel "example.com/go-gin-blog-api/comment/model"
	reactionModel "example.com/go-gin-blog-api/reaction/model"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Author   string `json:"author"`
	AuthorID uint   `json:"author_id"`

	User      model.User               `json:"user" gorm:"foreignKey:AuthorID;references:ID"`
	Comments  []commentModel.Comment   `json:"comments" gorm:"foreignKey:PostID"`
	Reactions []reactionModel.Reaction `json:"reactions" gorm:"foreignKey:PostID"`
}
