package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manuhdez/transactions-api/internal/accounts/controllers"
)

func main() {
	port := os.Getenv("APP_PORT")
	server := http.NewServeMux()
	server.HandleFunc("/status", controllers.StatusController)

	fmt.Printf("Transactions service running on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}
