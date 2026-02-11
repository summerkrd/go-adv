package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Db
	Email    string `json:"email"`
	Password string
	Address  string
}

type Db struct {
	DSN string
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		Db: Db{
			DSN: os.Getenv("DSN"),
		},
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Address:  os.Getenv("ADDRESS"),
	}
}
