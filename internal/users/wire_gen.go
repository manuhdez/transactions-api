// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/config"
	"github.com/manuhdez/transactions-api/internal/users/container"
	"github.com/manuhdez/transactions-api/internal/users/http/api"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/users/infra"
)

// Injectors from wire.go:

func Init() container.App {
	db := config.NewDBConnection()
	userMysqlRepository := infra.NewUserMysqlRepository(db)
	registerUser := service.NewRegisterUserService(userMysqlRepository)
	controllerRegisterUser := controller.NewRegisterUserController(registerUser)
	loginService := service.NewLoginService(userMysqlRepository)
	tokenService := infra.NewTokenService()
	login := controller.NewLoginController(loginService, tokenService)
	router := api.NewRouter(controllerRegisterUser, login)
	app := container.NewApp(router)
	return app
}
