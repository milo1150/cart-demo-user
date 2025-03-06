package api

import (
	"net/http"
	"user-service/internal/middlewares"
	"user-service/internal/repositories"
	"user-service/internal/schemas"
	"user-service/internal/services"
	"user-service/internal/types"
	"user-service/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func LoginHandler(c echo.Context, appState *types.AppState) error {
	payload := schemas.LoginPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	validate := validator.New()
	if errorMap := cartpkg.ValidateJsonPayload(validate, payload); errorMap != nil {
		return c.JSON(http.StatusBadRequest, errorMap)
	}

	// Verify username, password
	userService := services.UserService{DB: appState.DB}
	user, err := userService.VerifyUser(payload)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Create token with claims
	token, err := middlewares.CreateTokenWithClaims(*user, appState.Env.JwtSecret, appState.Env.JwtTokenDuration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Generate redis key
	key := utils.GenerateUserRedisKey(user.Username, c.RealIP())

	// Cache token in Redis
	if err := repositories.CacheUserToken(c, appState.RDB, key, token, appState.Env.JwtTokenDuration); err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Prepare response
	response := map[string]string{"token": token}

	return c.JSON(http.StatusOK, response)
}

func AuthHandler(c echo.Context, appState *types.AppState) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*middlewares.JwtCustomClaims)
	token := user.Raw
	redisKey := utils.GenerateUserRedisKey(claims.Name, c.RealIP())

	// Validate is token matched
	existedToken, err := repositories.FindUserToken(c, appState.RDB, redisKey)
	if token != existedToken || err != nil {
		return c.JSON(http.StatusUnauthorized, "expired jwt")
	}

	// Set forward header
	c.Response().Header().Set("X-User-Name", claims.Name)
	c.Response().Header().Set("X-User-Email", claims.Email)
	c.Response().Header().Set("X-User-Role", string(claims.Role))

	return c.NoContent(http.StatusOK)
}
