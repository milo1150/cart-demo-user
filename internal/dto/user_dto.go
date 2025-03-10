package dto

import "user-service/internal/models"

type UserInfo struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func TransformUserToUserInfo(user models.User) UserInfo {
	userInfo := UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
	return userInfo
}
