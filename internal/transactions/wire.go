//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/di"
)

func InitServer() di.Server {
	wire.Build(
		di.NewDBConnection,
		di.InitRepositories,
		di.InitBuses,
		di.InitServices,
		di.InitHandlers,
		di.InitControllers,
		di.NewServer,
	)

	return di.Server{}
}
