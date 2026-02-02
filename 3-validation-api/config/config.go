package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Email    string
	Password string
	Address  string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Address:  os.Getenv("ADDRESS"),
	}
}
