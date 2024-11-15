// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/manuhdez/transactions-api/internal/transactions/app/handler"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/container"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/transactions/http/router"
	"github.com/manuhdez/transactions-api/internal/transactions/infra"
	"github.com/manuhdez/transactions-api/shared/config"
)

// Injectors from wire.go:

func NewApp() container.App {
	dbConfig := config.NewDBConfig()
	db := config.NewDBConnection(dbConfig)
	transactionMysqlRepository := infra.NewTransactionMysqlRepository(db)
	accountMysqlRepository := infra.NewAccountMysqlRepository(db)
	createAccount := service.NewCreateAccountService(accountMysqlRepository)
	accountCreated := handler.NewAccountCreated(createAccount)
	eventBus := infra.NewEventBus(accountCreated)
	transactionService := service.NewTransactionService(transactionMysqlRepository, accountMysqlRepository, eventBus)
	deposit := controller.NewDeposit(transactionService)
	withdraw := controller.NewWithdraw(transactionService)
	findAllTransactions := service.NewFindAllTransactionsService(transactionMysqlRepository)
	controllerFindAllTransactions := controller.NewFindAllTransactions(findAllTransactions)
	findAccountTransactions := controller.NewFindAccountTransactions(transactionMysqlRepository)
	routerRouter := router.NewRouter(deposit, withdraw, controllerFindAllTransactions, findAccountTransactions)
	app := container.NewApp(routerRouter, eventBus)
	return app
}
