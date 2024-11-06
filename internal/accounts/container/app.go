package container

import (
	"github.com/google/wire"

	"github.com/manuhdez/transactions-api/internal/accounts/app/handler"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/accounts/http/router"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
	"github.com/manuhdez/transactions-api/shared/config"
)

var Databases = wire.NewSet(
	config.NewDBConfig,
	config.NewDBConnection,
)

var Repositories = wire.NewSet(
	wire.Bind(new(account.Repository), new(infra.AccountMysqlRepository)),
	infra.NewAccountMysqlRepository,
)

var Services = wire.NewSet(
	service.NewCreateService,
	service.NewFindAllService,
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
	wire.Bind(new(event.Bus), new(infra.EventBus)),
	infra.NewEventBus,
)

var EventHandlers = wire.NewSet(
	handler.NewHandlerDepositCreated,
	handler.NewWithdrawCreated,
)

type App struct {
	Server   router.Router
	EventBus event.Bus
}

func NewApp(
	router router.Router,
	eventBus event.Bus,
	depositCreatedHandler handler.DepositCreated,
	withdrawCreatedHandler handler.WithdrawCreated,
) App {
	eventBus.Subscribe(handler.DepositCreatedType, depositCreatedHandler)
	eventBus.Subscribe(handler.WithdrawCreatedType, withdrawCreatedHandler)

	return App{Server: router, EventBus: eventBus}
}
