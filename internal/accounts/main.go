package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/manuhdez/transactions-api/internal/accounts/bootstrap"
)

func main() {
	port := os.Getenv("APP_PORT")
	server := bootstrap.Server()

	fmt.Printf("Transactions service running on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), server))
}
