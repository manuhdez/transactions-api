package di

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/controllers"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
)

type Server struct {
	Engine   *gin.Engine
	EventBus event.Bus
}

func NewServer(
	eventBus event.Bus,
	statusController controllers.StatusController,
	depositController controllers.DepositController,
	findAllTransactionsController controllers.FindAllTransactionsController,
) Server {
	srv := gin.Default()

	// Register server routes
	srv.GET("/status", statusController.Handle)
	srv.POST("/deposit", depositController.Handle)
	srv.GET("/transactions", findAllTransactionsController.Handle)

	return Server{Engine: srv, EventBus: eventBus}
}
