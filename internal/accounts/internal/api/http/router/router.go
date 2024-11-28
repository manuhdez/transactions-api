package router

import (
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/api/http/v1/controller"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
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
	e.Validator = sharedhttp.NewRequestValidator()

	// Register global middleware
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/status", statusHandler)

	api := e.Group("/api")
	api.Use(echojwt.JWT([]byte(os.Getenv("JWT_SECRET"))))
	api.Use(sharedhttp.GetUserIdFromContext)

	v1 := api.Group("/v1")
	{
		v1.GET("/accounts", findAllAccounts.Handle)
		v1.POST("/accounts", createAccount.Handle)
		v1.GET("/accounts/:id", findAccount.Handle)
		v1.DELETE("/accounts/:id", deleteAccount.Handle)
	}

	e.GET("/metrics", sharedhttp.EchoWrapper(promhttp.Handler()))

	return Router{Engine: e}
}

func statusHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, echo.Map{"status": "ok", "service": "accounts"})
}
