//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/container"
)

func NewApp() container.App {
	wire.Build(
		container.Databases,
		container.Repositories,
		container.Services,
		container.Controllers,
		container.Router,
		container.Buses,
		container.EventHandlers,
		container.NewApp,
	)

	return container.App{}
}
