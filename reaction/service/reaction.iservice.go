package service

import (
	"example.com/go-gin-blog-api/reaction/model"
	"example.com/go-gin-blog-api/reaction/request"
)

type ReactionService interface {
	ReactToPost(username string, input request.ReactionInput) (*model.Reaction, error)
	GetReactionsByPost(postID string) ([]model.Reaction, error)
}
