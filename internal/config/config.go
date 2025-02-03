package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI string
	Port     string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load environment variables")
	}

	return &Config{
		MongoURI: os.Getenv("DATABASE_URL"),
		Port:     os.Getenv("PORT"),
	}
}
