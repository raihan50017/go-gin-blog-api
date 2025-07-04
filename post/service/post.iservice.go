package service

import "example.com/go-gin-blog-api/post/model"

type PostService interface {
	CreatePost(input model.Post, username string) (*model.Post, error)
	GetAllPosts() ([]model.Post, error)
	GetPostByID(id string) (*model.Post, error)
	UpdatePost(id string, username string, updated model.Post) (*model.Post, error)
	DeletePost(id string, username string) error
	GetUserPosts(username string) ([]model.Post, error)
}
