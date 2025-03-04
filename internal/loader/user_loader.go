package loader

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"user-service/internal/models"
	"user-service/internal/types"
	"user-service/internal/utils"

	"gorm.io/gorm"
)

func getUserJsonFile() []byte {
	basePath, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error get work directory")
	}

	userFilePath := filepath.Join(basePath, "internal", "assets", "user.json")
	bytes, err := os.ReadFile(userFilePath)
	if err != nil {
		log.Fatalf("Error read user.json: %v", err)
	}

	return bytes
}

func getDefaultUsers() types.UserJsonFile {
	userJsonFile := getUserJsonFile()
	users := &types.UserJsonFile{}

	if err := json.Unmarshal(userJsonFile, users); err != nil {
		log.Fatalf("Error parse user.json: %v", err)
	}

	return *users
}

func LoadDefaultUsers(db *gorm.DB) {
	users := getDefaultUsers()
	pwd := os.Getenv("ADMIN_PASSWORD")

	hashPassword, err := utils.HashPassword(pwd, 12)
	if err != nil {
		log.Fatalf("Error Hashing admin password")
	}

	for _, userJson := range users.Users {
		if err := db.First(&models.User{Username: userJson.Username}).Error; err != nil {
			newUser := &models.User{
				Username: userJson.Username,
				Password: hashPassword,
				Email:    userJson.Email,
			}
			if err := db.Create(newUser).Error; err != nil {
				log.Fatalf("Failed to create %v", userJson.Username)
			}
		}
	}
}
