package repository

import (
	authModel "example.com/go-gin-blog-api/auth/model"
	"example.com/go-gin-blog-api/post/model"
	"gorm.io/gorm"
)

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) Create(post *model.Post) error {
	return r.db.Create(post).Error
}

func (r *postRepository) FindAll() ([]model.Post, error) {
	var posts []model.Post
	err := r.db.
		Preload("User").
		Preload("Comments").
		Preload("Comments.User").
		Preload("Reactions").
		Preload("Reactions.User").
		Order("created_at desc").
		Find(&posts).Error
	return posts, err
}

func (r *postRepository) FindByID(id string) (*model.Post, error) {
	var post model.Post
	err := r.db.First(&post, id).Error
	return &post, err
}

func (r *postRepository) Update(post *model.Post) error {
	return r.db.Save(post).Error
}

func (r *postRepository) Delete(post *model.Post) error {
	return r.db.Delete(post).Error
}

func (r *postRepository) FindByAuthorID(authorID uint) ([]model.Post, error) {
	var posts []model.Post
	err := r.db.Where("author_id = ?", authorID).Order("created_at desc").Find(&posts).Error
	return posts, err
}

func (r *postRepository) FindUserByUsername(username string) (*authModel.User, error) {
	var user authModel.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}
