package services

import (
	"fmt"
	"user-service/internal/models"
	"user-service/internal/schemas"
	"user-service/internal/utils"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func (u *UserService) ValidateLoginUser(payload schemas.LoginPayload) error {
	// Find User
	user := models.User{}
	query := u.DB.First(&user, &models.User{Username: payload.Username})
	if err := query.Error; err != nil {
		return fmt.Errorf("user not found")
	}

	// Validate Password
	if ok := utils.CheckHashPassword(user.Password, payload.Password); !ok {
		return fmt.Errorf("invalid password")
	}

	return nil
}
