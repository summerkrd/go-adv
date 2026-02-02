package main

import (
	"go-adv/3-validation-api/config"
	"go-adv/3-validation-api/server"
	"go-adv/3-validation-api/verify"
)

func main() {
	conf := config.NewConfig()
	verifier := verify.NewVerifier(*conf)
	serv := server.NewServer()
}
