//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/di"
)

func NewServer() di.Server {
	wire.Build(
		di.Databases,
		di.Repositories,
		di.Services,
		di.Controllers,
		di.Buses,
		di.EventHandlers,
		di.Router,
		di.NewServer,
	)

	return di.Server{}
}
