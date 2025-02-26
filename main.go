package main

import (
	"net/http"
	"user-service/internal/database"
	"user-service/internal/loader"
	"user-service/internal/middlewares"
	"user-service/internal/routes"
	"user-service/internal/types"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load ENV
	loader.LoadEnv()

	// Database handler
	db := database.ConnectDatabase()
	database.RunAutoMigrate(db)

	// Global state
	appState := &types.AppState{
		DB: db,
	}

	// Creates an instance of Echo.
	e := echo.New()

	// Middlewares
	middlewares.RegisterMiddlewares(e)

	// Init Route
	routes.RegisterAppRoutes(e, appState)

	// TODO: delete
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "User service")
	})
	e.GET("/unauth", func(c echo.Context) error {
		return c.JSON(http.StatusUnauthorized, "User service - unauth")
	})
	// e.GET("/auth/validate", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, "Validate ok")
	// })
	e.GET("/user/login", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "User login")
	})
	e.GET("/user/validate", func(c echo.Context) error {
		// return c.JSON(http.StatusUnauthorized, "Unauth my man")
		return c.JSON(http.StatusOK, "You're good to go")
	})

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
