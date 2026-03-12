package main

import (
	"go-adv/4-order-api-start/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Product{}, &models.Session{}, &models.User{})
	if err != nil {
		panic(err)
	}
}
