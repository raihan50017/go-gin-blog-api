package models

import "gorm.io/gorm"

type Reaction struct {
	gorm.Model
	PostID   uint   `json:"post_id"`
	AuthorID uint   `json:"author_id"`
	Type     string `json:"type"`
}
