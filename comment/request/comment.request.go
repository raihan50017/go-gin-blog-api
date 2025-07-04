package request

type CommentInput struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}
