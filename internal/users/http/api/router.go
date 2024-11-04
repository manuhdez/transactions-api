package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/manuhdez/transactions-api/internal/users/http/api/middleware"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
)

type Router struct {
	port   string
	engine *mux.Router
}

func NewRouter(
	healthCheck controller.HealthCheck,
	registerUser controller.RegisterUser,
	loginUser controller.Login,
	getAllUsers controller.GetAllUsers,
) Router {
	router := mux.NewRouter()

	// middlewares
	router.Use(middleware.RequestLogger)
	router.Use(middleware.RequestMonitoring)

	// handlers
	router.HandleFunc("/health-check", healthCheck.Handle).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/auth/signup", registerUser.Handle).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/auth/login", loginUser.Handle).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/users", getAllUsers.Handle).Methods(http.MethodGet)

	// export prometheus metrics
	router.HandleFunc("/metrics", promhttp.Handler().ServeHTTP).Methods(http.MethodGet)

	return Router{
		port:   os.Getenv("APP_PORT"),
		engine: router,
	}
}

func (r Router) Listen() error {
	log.Print("Server running on", "port", r.port)
	return http.ListenAndServe(fmt.Sprintf(":%s", r.port), r.engine)
}
