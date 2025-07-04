package response

import (
	authResponse "example.com/go-gin-blog-api/auth/response"
	"example.com/go-gin-blog-api/reaction/model"
)

type ReactionResponse struct {
	ID        uint                      `json:"id"`
	Type      string                    `json:"type"`
	PostID    uint                      `json:"post_id"`
	Author    string                    `json:"author"`
	AuthorID  uint                      `json:"author_id"`
	CreatedAt string                    `json:"created_at"`
	User      authResponse.UserResponse `json:"user"`
}

func ToReactionResponse(reaction model.Reaction) ReactionResponse {
	return ReactionResponse{
		ID:        reaction.ID,
		Type:      reaction.Type,
		PostID:    reaction.PostID,
		Author:    reaction.Author,
		AuthorID:  reaction.AuthorID,
		CreatedAt: reaction.CreatedAt.Format("2006-01-02 15:04:05"),
		User:      authResponse.ToUserResponse(reaction.User),
	}
}

func ToReactionResponseList(comments []model.Reaction) []ReactionResponse {
	var result []ReactionResponse
	for _, c := range comments {
		result = append(result, ToReactionResponse(c))
	}
	return result
}
