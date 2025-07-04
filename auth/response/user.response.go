package response

import "example.com/go-gin-blog-api/auth/model"

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func ToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}
