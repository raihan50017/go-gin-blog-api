package service

import (
	"errors"

	"example.com/go-gin-blog-api/comment/model"
	"example.com/go-gin-blog-api/comment/repository"
	"example.com/go-gin-blog-api/comment/request"
	"example.com/go-gin-blog-api/comment/response"
)

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(repo repository.CommentRepository) CommentService {
	return &commentService{repo}
}

func (s *commentService) CreateComment(username string, input request.CommentInput) (*response.CommentResponse, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	comment := model.Comment{
		Content:  input.Content,
		PostID:   input.PostID,
		Author:   user.Username,
		AuthorID: user.ID,
	}

	if err := s.repo.Create(&comment); err != nil {
		return nil, errors.New("failed to create comment")
	}

	if err := s.repo.PreloadUser(&comment); err != nil {
		return nil, errors.New("failed to load comment user")
	}

	response := response.ToCommentResponse(comment)
	return &response, nil
}

func (s *commentService) GetCommentsByPost(postID string) ([]model.Comment, error) {
	return s.repo.GetByPostID(postID)
}
