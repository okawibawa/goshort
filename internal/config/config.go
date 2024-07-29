package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
		return nil, err
	}

	return &Config{
		DatabaseURL: os.Getenv("DB_URL"),
	}, nil
}
