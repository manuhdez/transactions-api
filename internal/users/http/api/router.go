package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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
) Router {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", healthCheck.Handle).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/auth/signup", registerUser.Handle).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/auth/login", loginUser.Handle).Methods(http.MethodPost)

	return Router{
		port:   os.Getenv("APP_PORT"),
		engine: router,
	}
}

func (r Router) Listen() error {
	fmt.Printf("Users service running on http://localhost:%s\n", r.port)
	return http.ListenAndServe(fmt.Sprintf(":%s", r.port), r.engine)
}
