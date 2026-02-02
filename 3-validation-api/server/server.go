package server

import (
	"net/http"
)

type Server struct {
	Router http.ServeMux
}

func NewServer() *Server {
	return &Server{
		Router: *http.NewServeMux(),
	}
}

func (s *Server) Start() error {

	err := http.ListenAndServe(":8080", &s.Router)
	if err != nil {
		return err
	}
	return nil
}
