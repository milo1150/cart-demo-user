package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid     uuid.UUID `gorm:"not null;uniqueIndex"`
	Username string    `gorm:"unique"`
	Name     string    `gorm:"unique"`
	Email    string    `gorm:"unique"`
	Password string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Uuid == uuid.Nil {
		uuidV7, err := uuid.NewV7()
		if err != nil {
			return err
		}
		u.Uuid = uuidV7
	}
	return nil
}
