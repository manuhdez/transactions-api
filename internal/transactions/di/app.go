package di

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/app/handler"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/config"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/transactions/http/router"
	"github.com/manuhdez/transactions-api/internal/transactions/infra"
)

var Databases = wire.NewSet(
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

var Buses = wire.NewSet(
	wire.Bind(new(event.Bus), new(infra.EventBus)),
	infra.NewEventBus,
)

var EventHandlers = wire.NewSet(
	handler.NewAccountCreated,
)

var Router = wire.NewSet(
	router.NewRouter,
)

type Server struct {
	Engine   *gin.Engine
	EventBus event.Bus
}

func NewServer(
	eventBus event.Bus,
	router router.Router,
	accountCreatedHandler handler.AccountCreated,
) Server {
	// Register event handlers
	eventBus.Subscribe(event.AccountCreatedType, accountCreatedHandler)

	return Server{
		Engine:   router.Engine,
		EventBus: eventBus,
	}
}
