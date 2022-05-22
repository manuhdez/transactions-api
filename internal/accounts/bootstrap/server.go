package bootstrap

import (
	"net/http"

	"github.com/manuhdez/transactions-api/internal/accounts/controllers"
)

func Server() http.Handler {
	deps := Deps()
	server := http.NewServeMux()

	server.HandleFunc("/status", controllers.StatusController)
	server.HandleFunc("/accounts", controllers.CreateAccountController(deps.Services.CreateAccount))
	return server
}
