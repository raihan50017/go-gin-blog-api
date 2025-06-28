package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content  string `json:"content" binding:"required"`
	PostID   uint   `json:"post_id"`
	Author   string `json:"author"`
	AuthorID uint   `json:"author_id"`
}
