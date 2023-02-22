// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/manuhdez/transactions-api/internal/transactions/app/handler"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/controllers"
	"github.com/manuhdez/transactions-api/internal/transactions/di"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/transactions/http/router"
	"github.com/manuhdez/transactions-api/internal/transactions/infra"
)

// Injectors from wire.go:

func InitServer() di.Server {
	eventBus := infra.NewEventBus()
	statusController := controllers.NewStatusController()
	db := di.NewDBConnection()
	transactionMysqlRepository := infra.NewTransactionMysqlRepository(db)
	deposit := service.NewDepositService(transactionMysqlRepository, eventBus)
	controllerDeposit := controller.NewDeposit(deposit)
	withdraw := service.NewWithdrawService(transactionMysqlRepository, eventBus)
	withdrawController := controllers.NewWithdrawController(withdraw)
	findAllTransactions := service.NewFindAllTransactionsService(transactionMysqlRepository)
	findAllTransactionsController := controllers.NewFindAllController(findAllTransactions)
	routerRouter := router.NewRouter(statusController, controllerDeposit, withdrawController, findAllTransactionsController)
	accountMysqlRepository := infra.NewAccountMysqlRepository(db)
	createAccount := service.NewCreateAccountService(accountMysqlRepository)
	accountCreated := handler.NewAccountCreated(createAccount)
	server := di.NewServer(eventBus, routerRouter, accountCreated)
	return server
}
