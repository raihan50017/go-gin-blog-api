package service

import (
	"errors"

	"example.com/go-gin-blog-api/post/model"
	"example.com/go-gin-blog-api/post/repository"
)

type postService struct {
	repo repository.PostRepository
}

func NewPostService(repo repository.PostRepository) PostService {
	return &postService{repo}
}

func (s *postService) CreatePost(input model.Post, username string) (*model.Post, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}
	input.Author = username
	input.AuthorID = user.ID
	if err := s.repo.Create(&input); err != nil {
		return nil, err
	}
	return &input, nil
}

func (s *postService) GetAllPosts() ([]model.Post, error) {
	return s.repo.FindAll()
}

func (s *postService) GetPostByID(id string) (*model.Post, error) {
	return s.repo.FindByID(id)
}

func (s *postService) UpdatePost(id string, username string, updated model.Post) (*model.Post, error) {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if post.Author != username {
		return nil, errors.New("unauthorized")
	}
	post.Title = updated.Title
	post.Content = updated.Content
	err = s.repo.Update(post)
	return post, err
}

func (s *postService) DeletePost(id string, username string) error {
	post, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if post.Author != username {
		return errors.New("unauthorized")
	}
	return s.repo.Delete(post)
}

func (s *postService) GetUserPosts(username string) ([]model.Post, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return s.repo.FindByAuthorID(user.ID)
}
