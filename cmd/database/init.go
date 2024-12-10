package database

import (
	"fmt"
	"github.com/hidalgo27/app-g1/cmd/utils/environment"
	"github.com/hidalgo27/app-g1/pkg/models"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

func LoadEnv() {
	env := os.Getenv("GO_ENV")
	var envFile string

	switch env {
	case "dev":
		envFile = ".env.dev"
	case "test":
		envFile = ".env.test"
	default:
		envFile = ".env"
	}

	fmt.Println(envFile)

	errorEnv := godotenv.Load(envFile)
	if errorEnv != nil {
		log.Fatal(errorEnv)
	}

	errVar := environment.ConfigVariables()
	if errVar != nil {
		log.Fatal(errVar)
	}
}

func CleanExpiredTokens() {
	err := DB.Where("expires_at < ?", time.Now()).Delete(&models.RefreshToken{}).Error
	if err != nil {
		log.Printf("Failed to clean expired tokens: %v", err)
	}
}
