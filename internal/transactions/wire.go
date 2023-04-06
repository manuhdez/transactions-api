//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/container"
)

func NewApp() container.App {
	wire.Build(
		container.Databases,
		container.Repositories,
		container.Services,
		container.Controllers,
		container.Buses,
		container.EventHandlers,
		container.Router,
		container.NewApp,
	)

	return container.App{}
}
