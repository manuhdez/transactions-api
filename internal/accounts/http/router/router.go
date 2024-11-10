package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/controller"
)

type Router struct {
	Engine *echo.Echo
}

func NewRouter(
	findAllAccounts controller.FindAllAccounts,
	createAccount controller.CreateAccount,
	findAccount controller.FindAccount,
	deleteAccount controller.DeleteAccount,
) Router {
	e := echo.New()

	// Register global middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/status", statusHandler)

	api := e.Group("/api")
	v1 := api.Group("/v1")
	{
		v1.GET("/accounts", findAllAccounts.Handle)
		v1.POST("/accounts", createAccount.Handle)
		v1.GET("/accounts/:id", findAccount.Handle)
		v1.DELETE("/accounts/:id", deleteAccount.Handle)
	}

	e.GET("/metrics", echoWrap(promhttp.Handler()))

	return Router{Engine: e}
}

func echoWrap(h http.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		h.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}

func statusHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"status": "ok", "service": "accounts"})
}
