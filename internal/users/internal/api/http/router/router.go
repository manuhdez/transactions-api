package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	controller2 "github.com/manuhdez/transactions-api/internal/users/internal/api/http/v1/controller"
	customMiddleware "github.com/manuhdez/transactions-api/internal/users/internal/api/http/v1/middleware"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
)

type Router struct {
	port   string
	engine *echo.Echo
}

func NewRouter(
	healthCheck controller2.HealthCheck,
	registerUser controller2.RegisterUser,
	loginUser controller2.Login,
	getAllUsers controller2.GetAllUsers,
) Router {
	e := echo.New()
	e.Validator = sharedhttp.NewRequestValidator()

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
