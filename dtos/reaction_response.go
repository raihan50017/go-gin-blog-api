package dtos

import "example.com/go-gin-blog-api/models"

type ReactionResponse struct {
	ID        uint         `json:"id"`
	Type      string       `json:"type"`
	PostID    uint         `json:"post_id"`
	Author    string       `json:"author"`
	AuthorID  uint         `json:"author_id"`
	CreatedAt string       `json:"created_at"`
	User      UserResponse `json:"user"`
}

func ToReactionResponse(reaction models.Reaction) ReactionResponse {
	return ReactionResponse{
		ID:        reaction.ID,
		Type:      reaction.Type,
		PostID:    reaction.PostID,
		Author:    reaction.Author,
		AuthorID:  reaction.AuthorID,
		CreatedAt: reaction.CreatedAt.Format("2006-01-02 15:04:05"),
		User:      ToUserResponse(reaction.User),
	}
}

func ToReactionResponseList(comments []models.Reaction) []ReactionResponse {
	var result []ReactionResponse
	for _, c := range comments {
		result = append(result, ToReactionResponse(c))
	}
	return result
}
