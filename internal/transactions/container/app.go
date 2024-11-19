package container

import (
	"github.com/google/wire"

	"github.com/manuhdez/transactions-api/shared/config"

	"github.com/manuhdez/transactions-api/internal/transactions/app/handler"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/transactions/http/router"
	"github.com/manuhdez/transactions-api/internal/transactions/infra"
)

var Databases = wire.NewSet(
	config.NewDBConfig,
	config.NewGormDBConnection,
)

var Repositories = wire.NewSet(
	wire.Bind(new(transaction.Repository), new(infra.TransactionMysqlRepository)),
	infra.NewTransactionMysqlRepository,
	wire.Bind(new(account.Repository), new(infra.AccountMysqlRepository)),
	infra.NewAccountMysqlRepository,
)
var Services = wire.NewSet(
	service.NewFindAllTransactionsService,
	service.NewCreateAccountService,
	service.NewTransactionService,
)

var Controllers = wire.NewSet(
	controller.NewDeposit,
	controller.NewWithdraw,
	controller.NewFindAllTransactions,
	controller.NewFindAccountTransactions,
	controller.NewTransferController,
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
