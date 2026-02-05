package server

import (
	"go-adv/3-validation-api/verify"
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

func (s *Server) RegisterRoutes(verifier *verify.Verifier) {
	s.Router.HandleFunc("POST /send", verifier.SendEmail)
	s.Router.HandleFunc("GET /verify/{hash}", verifier.VerifyHash)
}
