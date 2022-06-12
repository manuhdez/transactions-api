// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/bootstrap"
	"github.com/manuhdez/transactions-api/internal/accounts/controllers"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

// Injectors from wire.go:

func InitServer() bootstrap.Server {
	statusController := controllers.NewStatusController()
	db := bootstrap.InitializeDB()
	accountMysqlRepository := infra.NewAccountMysqlRepository(db)
	findAllService := service.NewFindAllService(accountMysqlRepository)
	findAllAccountsController := controllers.NewFindAllAccountsControllers(findAllService)
	createService := service.NewCreateService(accountMysqlRepository)
	createAccountController := controllers.NewCreateAccountController(createService)
	findAccountService := service.NewFindAccountService(accountMysqlRepository)
	findAccountController := controllers.NewFindAccountController(findAccountService)
	deleteAccountService := service.NewDeleteAccountService(accountMysqlRepository)
	deleteAccountController := controllers.NewDeleteAccountController(deleteAccountService)
	server := bootstrap.InitServer(statusController, findAllAccountsController, createAccountController, findAccountController, deleteAccountController)
	return server
}