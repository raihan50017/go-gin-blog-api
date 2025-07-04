package service

import (
	"errors"
	"time"

	"example.com/go-gin-blog-api/auth/model"
	"example.com/go-gin-blog-api/auth/repository"
	"example.com/go-gin-blog-api/utils"
	"golang.org/x/crypto/bcrypt"
)

const (
	accessTokenDuration  = 15 * time.Minute
	refreshTokenDuration = 24 * time.Hour
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo}
}

func (s *authService) Register(user *model.User) error {
	// Check duplicate email
	if existing, _ := s.repo.GetUserByEmail(user.Email); existing.ID != 0 {
		return errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("password hashing failed")
	}
	user.Password = string(hashed)

	return s.repo.CreateUser(user)
}

func (s *authService) Login(username, password string) (string, string, *model.User, error) {
	user, err := s.repo.GetUserByUsername(username)
	if err != nil {
		return "", "", nil, errors.New("invalid username or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", "", nil, errors.New("invalid username or password")
	}

	accessToken, _ := utils.GenerateToken(user.Username, accessTokenDuration)
	refreshToken, _ := utils.GenerateToken(user.Username, refreshTokenDuration)

	_ = s.repo.CreateRefreshToken(&model.RefreshToken{
		Token:     refreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(refreshTokenDuration).Unix(),
	})

	return accessToken, refreshToken, user, nil
}

func (s *authService) Refresh(oldToken string) (string, string, error) {
	claims, err := utils.ParseToken(oldToken)
	if err != nil {
		return "", "", errors.New("invalid refresh token")
	}

	stored, err := s.repo.GetRefreshToken(oldToken)
	if err != nil || time.Now().Unix() > stored.ExpiresAt {
		return "", "", errors.New("refresh token invalid or expired")
	}

	_ = s.repo.DeleteRefreshToken(oldToken)

	user, _ := s.repo.GetUserByUsername(claims.Username)
	newAccessToken, _ := utils.GenerateToken(claims.Username, accessTokenDuration)
	newRefreshToken, _ := utils.GenerateToken(claims.Username, refreshTokenDuration)

	_ = s.repo.CreateRefreshToken(&model.RefreshToken{
		Token:     newRefreshToken,
		UserID:    user.ID,
		ExpiresAt: time.Now().Add(refreshTokenDuration).Unix(),
	})

	return newAccessToken, newRefreshToken, nil
}
