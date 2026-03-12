package main

import (
	"go-adv/3-validation-api/config"
	"go-adv/4-order-api-start/auth"
	"go-adv/4-order-api-start/db"
	"go-adv/4-order-api-start/middleware"
	"go-adv/4-order-api-start/repository"
	"net/http"
)

func main() {
	conf := config.NewConfig()
	auth.InitJWT(conf.JWTSecret)
	dataBase := db.NewDb(conf)

	// Auth handlers
	authHandler := auth.NewAuthHandler(dataBase)

	router := http.NewServeMux()

	// Публичные роуты авторизации
	router.HandleFunc("POST /auth/send-code", authHandler.SendCode)
	router.HandleFunc("POST /auth/verify-code", authHandler.VerifyCode)

	// Защищенные роуты (требуют JWT)
	protectedRouter := http.NewServeMux()
	repository.NewRepoHandler(*protectedRouter, dataBase)
	router.Handle("/product/", middleware.AuthMiddleware(protectedRouter))

	handler := http.Handler(middleware.Log(router))
	http.ListenAndServe(":8081", handler)
}
