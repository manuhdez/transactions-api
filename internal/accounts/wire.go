//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/bootstrap"
)

func InitServer() bootstrap.Server {
	wire.Build(
		bootstrap.InitializeRepositories,
		bootstrap.InitBuses,
		bootstrap.InitServices,
		bootstrap.InitControllers,
		bootstrap.InitServer,
	)
	return bootstrap.Server{}
}
