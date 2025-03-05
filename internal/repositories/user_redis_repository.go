package repositories

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func CacheUserToken(c echo.Context, rdb *redis.Client, key string, token string, duration int) error {
	_, err := rdb.Set(c.Request().Context(), key, token, time.Hour*time.Duration(duration)).Result()
	if err != nil {
		return err
	}
	return nil
}

func FindUserToken(c echo.Context, rdb *redis.Client, key string) (string, error) {
	token, err := rdb.Get(c.Request().Context(), key).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}
