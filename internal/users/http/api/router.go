package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
)

type Router struct {
	Engine *mux.Router
}

func NewRouter(
	registerUser controller.RegisterUser,
	loginUser controller.Login,
) Router {
	router := mux.NewRouter()
	router.HandleFunc("/health-check", healthCheckController).Methods(http.MethodGet)
	router.HandleFunc("/api/v1/auth/signup", registerUser.Handle).Methods(http.MethodPost)
	router.HandleFunc("/api/v1/auth/login", loginUser.Handle).Methods(http.MethodPost)

	return Router{Engine: router}
}

func healthCheckController(w http.ResponseWriter, _ *http.Request) {
	type healthCheckResponse struct {
		Status  string `json:"status"`
		Service string `json:"service"`
	}

	body, _ := json.Marshal(healthCheckResponse{Status: "ok", Service: "users"})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}
