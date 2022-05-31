package di

import (
	"net/http"

	"github.com/manuhdez/transactions-api/internal/transactions/controllers"
)

func NewServer(
	statusController controllers.StatusController,
) *http.ServeMux {
	server := http.NewServeMux()

	// Register server routes
	server.HandleFunc("/status", statusController.Handle)

	return server
}
