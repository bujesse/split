package config

import (
	"log"
	"os"
	"split/config/logger"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	fxRatesApiToken := os.Getenv("SPLIT_FX_RATES_API_TOKEN")
	if fxRatesApiToken == "" {
		logger.Fatal("SPLIT_FX_RATES_API_TOKEN environment variable not set")
	}

	dbURL := os.Getenv("SPLIT_DATABASE_URL")
	if dbURL == "" {
		log.Fatal("SPLIT_DATABASE_URL is not set")
	}
	return err
}
