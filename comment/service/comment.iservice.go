package service

import (
	"example.com/go-gin-blog-api/comment/model"
	"example.com/go-gin-blog-api/comment/request"
	"example.com/go-gin-blog-api/comment/response"
)

type CommentService interface {
	CreateComment(username string, input request.CommentInput) (*response.CommentResponse, error)
	GetCommentsByPost(postID string) ([]model.Comment, error)
}
