package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/config"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/users/infra"
)

func main() {
	port := os.Getenv("APP_PORT")
	db := config.NewDBConnection()

	server := mux.NewRouter()
	server.HandleFunc("/status", StatusHandler).Methods(http.MethodGet)

	userRepository := infra.NewUserMysqlRepository(db)
	signupService := service.NewRegisterUserService(userRepository)
	signupController := controller.NewRegisterUserController(signupService)
	server.HandleFunc("/api/v1/auth/signup", signupController.Handle).Methods(http.MethodPost)

	fmt.Printf("Users service running on port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), server); err != nil {
		log.Fatalf("users service crashed: %e", err)
	}
}

func StatusHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("users service running ok"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
