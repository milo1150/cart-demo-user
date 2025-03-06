package main

import (
	"user-service/internal/database"
	"user-service/internal/loader"
	"user-service/internal/middlewares"
	"user-service/internal/nats"
	"user-service/internal/routes"
	"user-service/internal/types"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load ENV
	env := loader.LoadEnv()

	// NATS
	nc := nats.ConnectNATS()
	defer nc.Close()

	// Posrgres database
	db := database.ConnectPostgres()
	database.RunAutoMigrate(db)

	// Redis database
	rdb := database.ConnectRedis()
	defer rdb.Close()

	// Initialize Zap Logger
	logger := middlewares.InitializeZapLogger()

	// Global state
	appState := &types.AppState{
		DB:   db,
		RDB:  rdb,
		Env:  env,
		NATS: nc,
		Log:  logger,
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
