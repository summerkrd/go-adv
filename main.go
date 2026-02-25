package main

import (
	"go-adv/3-validation-api/config"
	"go-adv/3-validation-api/server"
	"go-adv/3-validation-api/verify"
	"go-adv/4-order-api-start/db"
	"go-adv/4-order-api-start/repository"
	"net/http"
)

func main() {
	conf := config.NewConfig()
	dataBase := db.NewDb(conf)
	verifier := verify.NewVerifier(*conf)
	serv := server.NewServer()
	serv.RegisterRoutes(verifier)
	repository.NewRepoHandler(serv.Router, dataBase)
	http.ListenAndServe(":8081", &serv.Router)
}
