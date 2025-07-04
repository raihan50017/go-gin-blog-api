package repository

import (
	authModel "example.com/go-gin-blog-api/auth/model"
	model "example.com/go-gin-blog-api/post/model"
)

type PostRepository interface {
	Create(post *model.Post) error
	FindAll() ([]model.Post, error)
	FindByID(id string) (*model.Post, error)
	Update(post *model.Post) error
	Delete(post *model.Post) error
	FindByAuthorID(authorID uint) ([]model.Post, error)
	FindUserByUsername(username string) (*authModel.User, error)
}
