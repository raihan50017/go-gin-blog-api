package repository

import (
	authModel "example.com/go-gin-blog-api/auth/model"
	"example.com/go-gin-blog-api/reaction/model"
)

type ReactionRepository interface {
	FindUserByUsername(username string) (*authModel.User, error)
	FindReactionByPostAndUser(postID uint, userID uint) (*model.Reaction, error)
	CreateReaction(reaction *model.Reaction) error
	UpdateReaction(reaction *model.Reaction) error
	GetReactionsByPost(postID string) ([]model.Reaction, error)
}
