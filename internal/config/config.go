package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	OpenExchangeRatesAPIKey string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("failed to load .env file")
	}

	apiKey := os.Getenv("OXR_API_KEY")
	if apiKey == "" {
		log.Fatal("OXR_API_KEY is not set (missing in env or .env file)")
	}

	return &Config{
		OpenExchangeRatesAPIKey: apiKey,
	}
}
