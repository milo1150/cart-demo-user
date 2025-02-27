package middlewares

import (
	"os"
	"user-service/internal/enums"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	IsAdmin bool       `json:"is_admin"`
	Role    enums.Role `json:"role"`
	jwt.RegisteredClaims
}

func JWT() echo.MiddlewareFunc {
	jwtSecret := os.Getenv("JWT_SECRET")

	return echojwt.WithConfig(echojwt.Config{
		SigningKey: jwtSecret,
	})
}
