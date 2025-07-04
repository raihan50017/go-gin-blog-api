package repository

import (
	authModel "example.com/go-gin-blog-api/auth/model"
	"example.com/go-gin-blog-api/reaction/model"
	"gorm.io/gorm"
)

type reactionRepo struct {
	db *gorm.DB
}

func NewReactionRepository(db *gorm.DB) ReactionRepository {
	return &reactionRepo{db}
}

func (r *reactionRepo) FindUserByUsername(username string) (*authModel.User, error) {
	var user authModel.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *reactionRepo) FindReactionByPostAndUser(postID uint, userID uint) (*model.Reaction, error) {
	var reaction model.Reaction
	err := r.db.Where("post_id = ? AND author_id = ?", postID, userID).First(&reaction).Error
	return &reaction, err
}

func (r *reactionRepo) CreateReaction(reaction *model.Reaction) error {
	return r.db.Create(reaction).Error
}

func (r *reactionRepo) UpdateReaction(reaction *model.Reaction) error {
	return r.db.Save(reaction).Error
}

func (r *reactionRepo) GetReactionsByPost(postID string) ([]model.Reaction, error) {
	var reactions []model.Reaction
	err := r.db.Where("post_id = ?", postID).Find(&reactions).Error
	return reactions, err
}
