// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/manuhdez/transactions-api/internal/accounts/app/handler"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/config"
	"github.com/manuhdez/transactions-api/internal/accounts/container"
	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/accounts/http/router"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

// Injectors from wire.go:

func NewApp() container.App {
	db := config.NewDBConnection()
	accountMysqlRepository := infra.NewAccountMysqlRepository(db)
	findAllService := service.NewFindAllService(accountMysqlRepository)
	findAllAccounts := controller.NewFindAllAccounts(findAllService)
	eventBus := infra.NewEventBus()
	createService := service.NewCreateService(accountMysqlRepository, eventBus)
	createAccount := controller.NewCreateAccount(createService)
	findAccountService := service.NewFindAccountService(accountMysqlRepository)
	findAccount := controller.NewFindAccountController(findAccountService)
	deleteAccountService := service.NewDeleteAccountService(accountMysqlRepository)
	deleteAccount := controller.NewDeleteAccount(deleteAccountService)
	routerRouter := router.NewRouter(findAllAccounts, createAccount, findAccount, deleteAccount)
	increaseBalanceService := service.NewIncreaseBalanceService(accountMysqlRepository)
	depositCreated := handler.NewHandlerDepositCreated(increaseBalanceService)
	decreaseBalance := service.NewDecreaseBalanceService(accountMysqlRepository)
	withdrawCreated := handler.NewWithdrawCreated(decreaseBalance)
	app := container.NewApp(routerRouter, eventBus, depositCreated, withdrawCreated)
	return app
}
