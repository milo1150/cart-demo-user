package loader

import (
	"log"
	"os"
	"strconv"
	"user-service/internal/types"

	"github.com/joho/godotenv"
)

func LoadEnv() types.Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize App Env
	env := types.Env{}
	env.JwtTokenDuration = GetJwtTokenDuration()
	env.JwtSecret = os.Getenv("JWT_SECRET")

	return env
}

func GetJwtTokenDuration() int {
	jwtTokenDurationEnv := os.Getenv("JWT_TOKEN_DURATION")

	jwtTokenDuration, err := strconv.Atoi(jwtTokenDurationEnv)
	if err != nil {
		log.Fatalf("Failed to load JWT_TOKEN_DURATION")
	}

	return jwtTokenDuration
}
