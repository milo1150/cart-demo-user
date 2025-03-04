package repositories

import (
	"user-service/internal/models"

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
