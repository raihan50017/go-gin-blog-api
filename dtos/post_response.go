package dtos

import "example.com/go-gin-blog-api/models"

type PostResponse struct {
	ID        uint               `json:"id"`
	Title     string             `json:"title"`
	Content   string             `json:"content"`
	Author    string             `json:"author"`
	AuthorID  uint               `json:"author_id"`
	CreatedAt string             `json:"created_at"`
	User      UserResponse       `json:"user"`
	Comments  []CommentResponse  `json:"comments"`
	Reactions []ReactionResponse `json:"reactions"`
}

func ToPostResponse(post models.Post) PostResponse {
	return PostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		Author:    post.Author,
		AuthorID:  post.AuthorID,
		CreatedAt: post.CreatedAt.Format("2006-01-02 15:04:05"),
		User:      ToUserResponse(post.User),
		Comments:  ToCommentResponseList(post.Comments),
		Reactions: ToReactionResponseList(post.Reactions),
	}
}
