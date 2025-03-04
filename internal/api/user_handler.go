package api

import (
	"net/http"
	"user-service/internal/middlewares"
	"user-service/internal/schemas"
	"user-service/internal/services"
	"user-service/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func LoginHandler(c echo.Context, appState types.AppState) error {
	payload := schemas.LoginPayload{}
	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	validate := validator.New()
	if errorMap := cartpkg.ValidateJsonPayload(validate, payload); errorMap != nil {
		return c.JSON(http.StatusBadRequest, errorMap)
	}

	// Validate username, password
	userService := services.UserService{DB: appState.DB}
	user, err := userService.ValidateLoginUser(payload)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Create token with claims
	env := appState.Env
	token, err := middlewares.CreateTokenWithClaims(*user, env.JwtSecret, env.JwtTokenDuration)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	response := map[string]string{"token": token}

	return c.JSON(http.StatusOK, response)
}
