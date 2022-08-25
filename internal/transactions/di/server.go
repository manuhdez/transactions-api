package di

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/controllers"
)

func NewServer(
	statusController controllers.StatusController,
	depositController controllers.DepositController,
	findAllTransactionsController controllers.FindAllTransactionsController,
) *gin.Engine {
	srv := gin.Default()

	// Register server routes
	srv.GET("/status", statusController.Handle)
	srv.POST("/deposit", depositController.Handle)
	srv.GET("/transactions", findAllTransactionsController.Handle)

	return srv
}
