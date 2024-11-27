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
	db := config.NewGormDBConnection(dbConfig)
	transactionMysqlRepository := infra.NewTransactionMysqlRepository(db)
	transactionService := service.NewTransactionService(transactionMysqlRepository)
	accountMysqlRepository := infra.NewAccountMysqlRepository(db)
	accountFinder := service.NewAccountFinder(accountMysqlRepository)
	createAccount := service.NewCreateAccountService(accountMysqlRepository)
	accountCreated := handler.NewAccountCreated(createAccount)
	eventBus := infra.NewEventBus(accountCreated)
	deposit := controller.NewDeposit(transactionService, accountFinder, eventBus)
	withdraw := controller.NewWithdraw(transactionService, accountFinder, eventBus)
	transfer := controller.NewTransferController(transactionService, accountFinder, eventBus)
	transactionsRetriever := service.NewTransactionsRetriever(transactionMysqlRepository)
	findAllTransactions := controller.NewFindAllTransactions(transactionsRetriever)
	findAccountTransactions := controller.NewFindAccountTransactions(transactionMysqlRepository)
	routerRouter := router.NewRouter(deposit, withdraw, transfer, findAllTransactions, findAccountTransactions)
	app := container.NewApp(routerRouter, eventBus)
	return app
}
