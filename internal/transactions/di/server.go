package di

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/handler"
	"github.com/manuhdez/transactions-api/internal/transactions/controllers"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
)

type Server struct {
	Engine   *gin.Engine
	EventBus event.Bus
}

func NewServer(
	eventBus event.Bus,
	accountCreatedHandler handler.AccountCreated,
	statusController controllers.StatusController,
	depositController controllers.DepositController,
	findAllTransactionsController controllers.FindAllTransactionsController,
	withdrawController controllers.WithdrawController,
) Server {
	srv := gin.Default()

	// Register server routes
	srv.GET("/status", statusController.Handle)
	srv.POST("/deposit", depositController.Handle)
	srv.POST("/withdraw", withdrawController.Handle)
	srv.GET("/transactions", findAllTransactionsController.Handle)

	// Register event handlers
	eventBus.Subscribe(event.AccountCreatedType, accountCreatedHandler)

	return Server{Engine: srv, EventBus: eventBus}
}
