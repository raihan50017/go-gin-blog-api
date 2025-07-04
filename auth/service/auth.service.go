package service

import "example.com/go-gin-blog-api/auth/model"

type AuthService interface {
	Register(user *model.User) error
	Login(username, password string) (string, string, *model.User, error)
	Refresh(oldToken string) (string, string, error)
}
