package router

import (
	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/controllers"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
)

type Router struct {
	Engine *gin.Engine
}

// NewRouter Router initialiser for dependency injection. Returns a gin.Engine instance.
func NewRouter(
	statusController controllers.StatusController,
	depositController controller.Deposit,
	withdrawController controllers.WithdrawController,
	findAllTransactionsController controllers.FindAllTransactionsController,
) Router {
	router := gin.Default()

	router.GET("/status", statusController.Handle)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/deposit", depositController.Handle)
			v1.POST("/withdraw", withdrawController.Handle)
			v1.GET("/transactions", findAllTransactionsController.Handle)
		}
	}

	return Router{router}
}
