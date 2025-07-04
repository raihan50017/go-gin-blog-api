package repository

import (
	authModel "example.com/go-gin-blog-api/auth/model"
	model "example.com/go-gin-blog-api/comment/model"
)

type CommentRepository interface {
	Create(comment *model.Comment) error
	GetByPostID(postID string) ([]model.Comment, error)
	FindUserByUsername(username string) (*authModel.User, error)
	PreloadUser(comment *model.Comment) error
}
