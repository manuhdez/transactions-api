package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	customMiddleware "github.com/manuhdez/transactions-api/internal/users/http/api/middleware"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
)

type Router struct {
	port   string
	engine *echo.Echo
}

func NewRouter(
	healthCheck controller.HealthCheck,
	registerUser controller.RegisterUser,
	loginUser controller.Login,
	getAllUsers controller.GetAllUsers,
) Router {
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(customMiddleware.RequestMonitoring)

	// handlers
	e.GET("/health-check", healthCheck.Handle)

	v1 := e.Group("/api/v1")
	{
		v1.POST("/auth/signup", registerUser.Handle)
		v1.POST("/auth/login", loginUser.Handle)
		v1.GET("/users", getAllUsers.Handle)
	}

	// export prometheus metrics
	e.GET("/metrics", sharedhttp.EchoWrapper(promhttp.Handler()))

	return Router{
		port:   os.Getenv("APP_PORT"),
		engine: e,
	}
}

func (r Router) Listen() error {
	log.Print("Server running on", "port", r.port)
	return http.ListenAndServe(fmt.Sprintf(":%s", r.port), r.engine)
}
