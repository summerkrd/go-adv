package db

import (
	"go-adv/3-validation-api/config"
	"go-adv/4-order-api-start/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *config.Config) *Db {
	db, err := gorm.Open(postgres.Open(conf.Db.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&models.Product{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
