package repository

import (
	"example.com/go-gin-blog-api/auth/model"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db}
}

func (r *authRepository) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *authRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *authRepository) CreateUser(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *authRepository) CreateRefreshToken(token *model.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *authRepository) GetRefreshToken(token string) (*model.RefreshToken, error) {
	var rt model.RefreshToken
	err := r.db.Where("token = ?", token).First(&rt).Error
	return &rt, err
}

func (r *authRepository) DeleteRefreshToken(token string) error {
	return r.db.Where("token = ?", token).Delete(&model.RefreshToken{}).Error
}
