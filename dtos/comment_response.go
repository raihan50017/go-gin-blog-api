package dtos

import "example.com/go-gin-blog-api/models"

type CommentResponse struct {
	ID        uint         `json:"id"`
	Content   string       `json:"content"`
	PostID    uint         `json:"post_id"`
	Author    string       `json:"author"`
	AuthorID  uint         `json:"author_id"`
	CreatedAt string       `json:"created_at"`
	User      UserResponse `json:"user"`
}

func ToCommentResponse(comment models.Comment) CommentResponse {
	return CommentResponse{
		ID:        comment.ID,
		Content:   comment.Content,
		PostID:    comment.PostID,
		Author:    comment.Author,
		AuthorID:  comment.AuthorID,
		CreatedAt: comment.CreatedAt.Format("2006-01-02 15:04:05"),
		User:      ToUserResponse(comment.User),
	}
}

func ToCommentResponseList(comments []models.Comment) []CommentResponse {
	var result []CommentResponse
	for _, c := range comments {
		result = append(result, ToCommentResponse(c))
	}
	return result
}
