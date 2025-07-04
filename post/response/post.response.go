package response

import (
	authResponse "example.com/go-gin-blog-api/auth/response"
	commentResponse "example.com/go-gin-blog-api/comment/response"
	"example.com/go-gin-blog-api/post/model"
	reactionResponse "example.com/go-gin-blog-api/reaction/response"
)

type PostResponse struct {
	ID        uint                                `json:"id"`
	Title     string                              `json:"title"`
	Content   string                              `json:"content"`
	Author    string                              `json:"author"`
	AuthorID  uint                                `json:"author_id"`
	CreatedAt string                              `json:"created_at"`
	User      authResponse.UserResponse           `json:"user"`
	Comments  []commentResponse.CommentResponse   `json:"comments"`
	Reactions []reactionResponse.ReactionResponse `json:"reactions"`
}

func ToPostResponse(post model.Post) PostResponse {
	return PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Author:    post.Author,
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		User:      authResponse.ToUserResponse(post.User),
		Comments:  commentResponse.ToCommentResponseList(post.Comments),
		Reactions: reactionResponse.ToReactionResponseList(post.Reactions),
	}
}
