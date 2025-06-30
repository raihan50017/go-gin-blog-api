package models

import "gorm.io/gorm"

type Reaction struct {
	gorm.Model
	PostID   uint   `json:"post_id"`
	AuthorID uint   `json:"author_id"`
	Author   string `json:"author"`
	Type     string `json:"type"`

	Post Post `json:"post" gorm:"foreignKey:PostID;references:ID"`
	User User `json:"user" gorm:"foreignKey:AuthorID;references:ID"`
}
