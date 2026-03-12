package repository

import (
	"encoding/json"
	"go-adv/4-order-api-start/db"
	"net/http"
	"strconv"
)

func NewRepoHandler(router http.ServeMux, db *db.Db) {
	repo := NewRepository(db)
	router.HandleFunc("POST /product/create", CreateProduct(repo))
	router.HandleFunc("POST /product/update", UpdateProduct(repo))
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
		product := Product{}
		json.NewDecoder(r.Body).Decode(&product)
		repo.Update(&product)
	}
}

func DeleteProduct(repo *Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		idInt, _ := strconv.Atoi(idStr)
		repo.Delete(uint(idInt))
	}
}

func GetByID(repo *Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		idInt, _ := strconv.Atoi(idStr)
		product, _ := repo.GetByID(uint(idInt))
		json.NewEncoder(w).Encode(product)
	}
}
