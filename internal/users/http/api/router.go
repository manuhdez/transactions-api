package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
)

type Router struct {
	Engine *mux.Router
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

	return Router{Engine: router}
}
