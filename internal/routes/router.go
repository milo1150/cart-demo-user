package routes

import (
	"net/http"
	"user-service/internal/middlewares"
	"user-service/internal/types"

	"github.com/labstack/echo/v4"
)

type RegisterRoutes struct {
	Echo     *echo.Echo
	AppState *types.AppState
}

func (r *RegisterRoutes) RegisterAppRoutes() {
	r.publicRoutes()
	r.privateRoutes()
}

func (r *RegisterRoutes) publicRoutes() {
	r.Echo.POST("/login", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/login endpoint")
	})
}

func (r *RegisterRoutes) privateRoutes() {
	userGroup := r.Echo.Group("/user")
	userGroup.Use(middlewares.JWT())

	userGroup.POST("/create", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "/create user endpoint")
	})

	userGroup.GET("/auth", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "auth endpoint")
	})
}
