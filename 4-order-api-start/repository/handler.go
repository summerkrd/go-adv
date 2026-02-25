package repository

import (
	"encoding/json"
	"go-adv/4-order-api-start/db"
	"net/http"
)

func NewRepoHandler(router http.ServeMux, db *db.Db) {
	repo := NewRepository(db)
	router.HandleFunc("POST /product/create", CreateProduct(repo))
	router.HandleFunc("POST /product/update/{id}", UpdateProduct(repo))
	router.HandleFunc("POST /product/delete/{id}", DeleteProduct(repo))
	router.HandleFunc("GET /product/{id}", GetByID(repo))
}

func CreateProduct(repo *Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var product Product
		json.NewDecoder(r.Body).Decode(&product)
		repo.Create(&product)
	}
}

func UpdateProduct(repo *Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func DeleteProduct(repo *Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func GetByID(repo *Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
