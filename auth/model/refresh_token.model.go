package model

import "gorm.io/gorm"

type RefreshToken struct {
	gorm.Model
	Token     string `gorm:"uniqueIndex"`
	UserID    uint
	ExpiresAt int64
}
