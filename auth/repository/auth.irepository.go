package repository

import "example.com/go-gin-blog-api/auth/model"

type AuthRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	GetUserByUsername(username string) (*model.User, error)
	CreateUser(user *model.User) error
	CreateRefreshToken(token *model.RefreshToken) error
	GetRefreshToken(token string) (*model.RefreshToken, error)
	DeleteRefreshToken(token string) error
}
