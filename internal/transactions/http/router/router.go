package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
)

type Router struct {
	Engine *gin.Engine
}

// NewRouter initializes a new router for dependency injection, returning a gin.Engine instance.
func NewRouter(
	depositController controller.Deposit,
	withdrawController controller.Withdraw,
	findAllTransactionsController controller.FindAllTransactions,
	findAccountTransactions controller.FindAccountTransactions,
) Router {
	router := gin.Default()

	router.GET("/status", statusHandler)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.POST("/deposit", depositController.Handle)
			v1.POST("/withdraw", withdrawController.Handle)
			v1.GET("/transactions", findAllTransactionsController.Handle)
			v1.GET("/transactions/:id", findAccountTransactions.Handle)
		}
	}

	return Router{router}
}

func statusHandler(ctx *gin.Context) {
	type statusResponse struct {
		Status  string `json:"status"`
		Service string `json:"service"`
	}

	ctx.JSON(http.StatusOK, statusResponse{"ok", "transactions"})
}
