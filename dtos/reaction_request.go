package dtos

type ReactionInput struct {
	Type   string `json:"type" binding:"required"`
	PostID uint   `json:"post_id" binding:"required"`
}
