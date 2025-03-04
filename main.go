package main

import (
	"user-service/internal/database"
	"user-service/internal/loader"
	"user-service/internal/middlewares"
	"user-service/internal/routes"
	"user-service/internal/types"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load ENV
	env := loader.LoadEnv()

	// Database handler
	db := database.ConnectDatabase()
	database.RunAutoMigrate(db)

	// Global state
	appState := &types.AppState{
		DB:  db,
		Env: env,
	}

	// Initialize User if run first time
	loader.LoadDefaultUsers(db)

	// Creates an instance of Echo.
	e := echo.New()

	// Middlewares
	middlewares.RegisterMiddlewares(e)

	// Init Route
	r := routes.RegisterRoutes{Echo: e, AppState: appState}
	r.RegisterAppRoutes()

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
