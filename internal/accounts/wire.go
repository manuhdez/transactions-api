//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/bootstrap"
	"github.com/manuhdez/transactions-api/internal/accounts/http/router"
)

func InitServer() bootstrap.Server {
	wire.Build(
		bootstrap.InitializeRepositories,
		bootstrap.InitBuses,
		bootstrap.InitServices,
		bootstrap.InitHandlers,
		bootstrap.InitControllers,
		router.NewRouter,
		bootstrap.InitServer,
	)
	return bootstrap.Server{}
}
