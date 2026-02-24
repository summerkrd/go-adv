package repository

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	Router http.ServeMux
}

func NewServer() *Handler {
	return &Handler{
		Router: *http.NewServeMux(),
	}
}

func (s *Handler) CreateProduct(repository Repository) {
	s.Router.HandleFunc("POST /create", func(writer http.ResponseWriter, request *http.Request) {
		var product Product
		json.NewDecoder(request.Body).Decode(&product)
		repository.Create(&product)
	})

}
