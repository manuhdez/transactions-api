package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
)

type Router struct {
	Engine *echo.Echo
}

// NewRouter initializes a new router for dependency injection, returning a gin.Engine instance.
func NewRouter(
	depositController controller.Deposit,
	withdrawController controller.Withdraw,
	findAllTransactionsController controller.FindAllTransactions,
	findAccountTransactions controller.FindAccountTransactions,
) Router {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/status", statusHandler)

	api := e.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.POST("/deposit", depositController.Handle)
		v1.POST("/withdraw", withdrawController.Handle)
		v1.GET("/transactions", findAllTransactionsController.Handle)
		v1.GET("/transactions/:id", findAccountTransactions.Handle)
	}

	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return Router{Engine: e}
}

func statusHandler(c echo.Context) error {
	return c.JSON(
		http.StatusOK,
		echo.Map{"status": "ok", "service": "transactions"},
	)
}
