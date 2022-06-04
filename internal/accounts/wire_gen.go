// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/manuhdez/transactions-api/internal/accounts/bootstrap"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

// Injectors from wire.go:

func InitializeServer() bootstrap.Server {
	db := bootstrap.InitializeDB()
	accountMysqlRepository := infra.NewAccountMysqlRepository(db)
	services := bootstrap.InitializeServices(accountMysqlRepository)
	controllers := bootstrap.InitializeControllers(services)
	server := bootstrap.InitializeServer(controllers)
	return server
}
