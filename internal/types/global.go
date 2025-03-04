package types

import (
	"gorm.io/gorm"
)

type Env struct {
	JwtTokenDuration int
	JwtSecret        string
}

type AppState struct {
	DB  *gorm.DB
	Env Env
}
