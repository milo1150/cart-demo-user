package types

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Env struct {
	JwtTokenDuration int
	JwtSecret        string
}

type AppState struct {
	DB  *gorm.DB
	RDB *redis.Client
	Env Env
}
