package main

import (
	"go-adv/3-validation-api/config"
	"go-adv/3-validation-api/server"
	"go-adv/3-validation-api/verify"
	"go-adv/4-order-api-start/auth"
	"go-adv/4-order-api-start/db"
	"go-adv/4-order-api-start/middleware"
	"go-adv/4-order-api-start/repository"
	"net/http"
)

func main() {
	conf := config.NewConfig()
	dataBase := db.NewDb(conf)

	// Auth handlers
	authHandler := auth.NewAuthHandler(dataBase)

	verifier := verify.NewVerifier(*conf)
	serv := server.NewServer()
	serv.RegisterRoutes(verifier)

	// Публичные роуты авторизации
	serv.Router.HandleFunc("POST /auth/send-code", authHandler.SendCode)
	serv.Router.HandleFunc("POST /auth/verify-code", authHandler.VerifyCode)

	// Защищенные роуты (требуют JWT)
	protectedRouter := http.NewServeMux()
	repository.NewRepoHandler(*protectedRouter, dataBase)
	serv.Router.Handle("/product/", middleware.AuthMiddleware(protectedRouter))

	handler := http.Handler(middleware.Log(&serv.Router))
	http.ListenAndServe(":8081", handler)

	//repository.NewRepoHandler(serv.Router, dataBase)
	//handler := http.Handler(middleware.Log(&serv.Router))
	//http.ListenAndServe(":8081", handler)
}
