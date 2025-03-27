package middlewares

import (
	"os"
	"time"
	"user-service/internal/enums"
	"user-service/internal/models"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	UserId  uint       `json:"user_id"`
	Name    string     `json:"name"`
	Email   string     `json:"email"`
	IsAdmin bool       `json:"is_admin"`
	Role    enums.Role `json:"role"`
	jwt.RegisteredClaims
}

func JwtMiddleware() echo.MiddlewareFunc {
	jwtSecret := os.Getenv("JWT_SECRET")

	// Validate jwt
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(jwtSecret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
	})
}

func CreateTokenWithClaims(user models.User, secret string, tokenDuration int) (string, error) {
	claims := JwtCustomClaims{
		UserId: user.ID,
		Name:   user.Username,
		Email:  user.Email,
		Role:   enums.Admin, // TODO: user.Role
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(tokenDuration))),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := jwtToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, nil
}
