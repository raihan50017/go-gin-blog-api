package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Author   string `json:"author"`
	AuthorID uint   `json:"author_id"`

	User      User       `json:"user" gorm:"foreignKey:AuthorID;references:ID"`
	Comments  []Comment  `json:"comments" gorm:"foreignKey:PostID"`
	Reactions []Reaction `json:"reactions" gorm:"foreignKey:PostID"`
}
