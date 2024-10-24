package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

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
	// define a middleware function to log the request method, path and request body
	loggerMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Printf("%s %s %s", r.Method, r.URL.Path, r.Proto)
			next.ServeHTTP(w, r)
		})
	}
	// add the middleware to the router
	router.Use(loggerMiddleware)
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
