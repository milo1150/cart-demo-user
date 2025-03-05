package routes

import (
	"net/http"
	"user-service/internal/api"
	"user-service/internal/middlewares"
	"user-service/internal/types"

	"github.com/labstack/echo/v4"
)

type RegisterRoutes struct {
	Echo     *echo.Echo
	AppState *types.AppState
}

func (r *RegisterRoutes) RegisterAppRoutes() {
	userGroup := r.Echo.Group("/user")
	r.publicRoutes(*userGroup)
	r.privateRoutes(*userGroup)
}

func (r *RegisterRoutes) publicRoutes(userGroup echo.Group) {
	userGroup.POST("/login", func(c echo.Context) error {
		return api.LoginHandler(c, r.AppState)
	})
}

func (r *RegisterRoutes) privateRoutes(userGroup echo.Group) {
	// If not implment jwt config middlware, in "user" context in AuthHandler wil be nil
	userGroup.Use(middlewares.JwtMiddleware())

	// Gateway Forward auth only.
	// Forward auth can't debug.
	userGroup.GET("/auth", func(c echo.Context) error {
		return api.AuthHandler(c, r.AppState)
	})

	userGroup.POST("/create", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/create user endpoint") // TODO
	})
}
