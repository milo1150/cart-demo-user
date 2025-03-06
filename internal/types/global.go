package types

import (
	"github.com/nats-io/nats.go"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Env struct {
	JwtTokenDuration int
	JwtSecret        string
}

type AppState struct {
	DB   *gorm.DB
	RDB  *redis.Client
	Env  Env
	NATS *nats.Conn
	Log  *zap.Logger
}
