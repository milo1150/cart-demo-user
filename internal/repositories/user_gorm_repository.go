package repositories

import (
	"user-service/internal/models"
	"user-service/internal/schemas"
	"user-service/internal/utils"

	"gorm.io/gorm"
)

func FindUser(db *gorm.DB, username string) (*models.User, error) {
	user := &models.User{}
	query := db.First(user, &models.User{Username: username})
	if err := query.Error; err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUser(db *gorm.DB, payload schemas.CreateUserPayload) (*models.User, error) {
	encryptPassword, err := utils.HashPassword(payload.Password, 10)
	if err != nil {
		return nil, err
	}

	createUser := models.User{
		Username: payload.Username,
		Password: encryptPassword,
		Name:     payload.Name,
		Email:    payload.Email,
	}

	if err := db.Create(&createUser).Error; err != nil {
		return nil, err
	}

	return &createUser, nil
}
