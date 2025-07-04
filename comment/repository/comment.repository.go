package repository

import (
	authModel "example.com/go-gin-blog-api/auth/model"
	"example.com/go-gin-blog-api/comment/model"
	"gorm.io/gorm"
)

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepo{db}
}

func (r *commentRepo) Create(comment *model.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepo) GetByPostID(postID string) ([]model.Comment, error) {
	var comments []model.Comment
	err := r.db.Where("post_id = ?", postID).Order("created_at").Find(&comments).Error
	return comments, err
}

func (r *commentRepo) FindUserByUsername(username string) (*authModel.User, error) {
	var user authModel.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *commentRepo) PreloadUser(comment *model.Comment) error {
	return r.db.Preload("User").First(comment, comment.ID).Error
}
