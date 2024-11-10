package container

import (
	"github.com/google/wire"

	"github.com/manuhdez/transactions-api/internal/transactions/app/handler"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/transactions/http/router"
	"github.com/manuhdez/transactions-api/internal/transactions/infra"
	"github.com/manuhdez/transactions-api/shared/config"
)

var Databases = wire.NewSet(
	config.NewDBConfig,
	config.NewDBConnection,
)

var Repositories = wire.NewSet(
	wire.Bind(new(transaction.Repository), new(infra.TransactionMysqlRepository)),
	infra.NewTransactionMysqlRepository,
	wire.Bind(new(account.Repository), new(infra.AccountMysqlRepository)),
	infra.NewAccountMysqlRepository,
)
var Services = wire.NewSet(
	service.NewDepositService,
	service.NewFindAllTransactionsService,
	service.NewCreateAccountService,
	service.NewWithdrawService,
)

var Controllers = wire.NewSet(
	controller.NewDeposit,
	controller.NewWithdraw,
	controller.NewFindAllTransactions,
	controller.NewFindAccountTransactions,
)

var EventHandlers = wire.NewSet(
	handler.NewAccountCreated,
)

var Buses = wire.NewSet(
	wire.Bind(new(event.Bus), new(infra.EventBus)),
	infra.NewEventBus,
)

var Router = wire.NewSet(
	router.NewRouter,
)

type App struct {
	Server   router.Router
	EventBus event.Bus
}

func NewApp(router router.Router, eventBus event.Bus) App {
	return App{
		Server:   router,
		EventBus: eventBus,
	}
}
