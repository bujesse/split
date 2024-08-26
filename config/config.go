package config

import (
	"os"
	"split/config/logger"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	fxRatesApiToken := GetFxRatesApiToken()
	if fxRatesApiToken == "" {
		logger.Fatal("SPLIT_FX_RATES_API_TOKEN environment variable not set")
	}
	return err
}

func GetFxRatesApiToken() string {
	return os.Getenv("SPLIT_FX_RATES_API_TOKEN")
}
