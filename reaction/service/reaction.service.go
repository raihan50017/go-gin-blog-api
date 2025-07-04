package service

import (
	"errors"

	"example.com/go-gin-blog-api/reaction/model"
	"example.com/go-gin-blog-api/reaction/repository"
	"example.com/go-gin-blog-api/reaction/request"
)

type reactionService struct {
	repo repository.ReactionRepository
}

func NewReactionService(repo repository.ReactionRepository) ReactionService {
	return &reactionService{repo}
}

func (s *reactionService) ReactToPost(username string, input request.ReactionInput) (*model.Reaction, error) {
	user, err := s.repo.FindUserByUsername(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	existing, err := s.repo.FindReactionByPostAndUser(input.PostID, user.ID)
	if err == nil {
		existing.Type = input.Type
		if err := s.repo.UpdateReaction(existing); err != nil {
			return nil, err
		}
		return existing, nil
	}

	reaction := &model.Reaction{
		Type:     input.Type,
		PostID:   input.PostID,
		Author:   user.Username,
		AuthorID: user.ID,
	}
	if err := s.repo.CreateReaction(reaction); err != nil {
		return nil, err
	}

	return reaction, nil
}

func (s *reactionService) GetReactionsByPost(postID string) ([]model.Reaction, error) {
	return s.repo.GetReactionsByPost(postID)
}
