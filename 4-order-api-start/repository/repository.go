package repository

import (
	"go-adv/4-order-api-start/db"
	"go-adv/4-order-api-start/models"

	"gorm.io/gorm/clause"
)

type Product = models.Product

type Repository struct {
	Database *db.Db
}

func NewRepository(database *db.Db) *Repository {
	return &Repository{
		Database: database,
	}
}

func (repo *Repository) Create(product *Product) (*Product, error) {
	result := repo.Database.DB.Create(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *Repository) Update(product *Product) (*Product, error) {
	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(product)
	if result.Error != nil {
		return nil, result.Error
	}
	return product, nil
}

func (repo *Repository) Delete(id uint) error {
	result := repo.Database.DB.Delete(&Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *Repository) GetByID(id uint) (*Product, error) {
	var product Product
	result := repo.Database.DB.First(&product, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &product, nil
}
