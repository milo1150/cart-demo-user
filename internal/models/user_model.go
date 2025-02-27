package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uuid     uuid.UUID `gorm:"not null;uniqueIndex"`
	Username string
	Email    string
	Password string
	Name     string
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
