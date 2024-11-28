//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/api/http/router"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/api/http/v1/controller"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/event_handler"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/infra/db"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/infra/queue/rabbitmq"
	"github.com/manuhdez/transactions-api/shared/config"
)

var Databases = wire.NewSet(
	config.NewDBConfig,
	config.NewGormDBConnection,
)

var Repositories = wire.NewSet(
	wire.Bind(new(account.Repository), new(db.AccountPostgresRepository)),
	db.NewAccountPostgresRepository,
)

var Services = wire.NewSet(
	service.NewCreateService,
	service.NewFindAccountService,
	service.NewDeleteAccountService,
	service.NewIncreaseBalanceService,
	service.NewDecreaseBalanceService,
)

var Controllers = wire.NewSet(
	controller.NewCreateAccount,
	controller.NewFindAccountController,
	controller.NewFindAllAccounts,
	controller.NewDeleteAccount,
)

var Router = wire.NewSet(router.NewRouter)

var Buses = wire.NewSet(
	wire.Bind(new(event.Bus), new(rabbitmq.EventBus)),
	rabbitmq.NewEventBus,
)

var EventHandlers = wire.NewSet(
	event_handler.NewHandlerDepositCreated,
	event_handler.NewWithdrawCreated,
)

type App struct {
	Server   router.Router
	EventBus event.Bus
}

// TODO: move event subscription to a proper place
func NewApp(
	router router.Router,
	eventBus event.Bus,
	depositCreatedHandler event_handler.DepositCreated,
	withdrawCreatedHandler event_handler.WithdrawCreated,
) App {
	eventBus.Subscribe(event_handler.DepositCreatedType, depositCreatedHandler)
	eventBus.Subscribe(event_handler.WithdrawCreatedType, withdrawCreatedHandler)

	return App{Server: router, EventBus: eventBus}
}

func BootstrapApp() App {
	wire.Build(Databases, Repositories, Services, Controllers, Router, Buses, EventHandlers, NewApp)
	return App{}
}
