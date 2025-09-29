package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TelegramToken string
	PostgresDSN   string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, relying on environment variables")
	}

	return &Config{
		TelegramToken: os.Getenv("TELEGRAM_TOKEN"),
		PostgresDSN:   os.Getenv("POSTGRES_DSN"),
	}
}
