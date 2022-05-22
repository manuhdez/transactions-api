package main

import (
	"fmt"
	"log"
	"os"

	"github.com/manuhdez/transactions-api/internal/accounts/bootstrap"
)

func main() {
	port := os.Getenv("APP_PORT")
	server := bootstrap.Server()

	fmt.Printf("Transactions service running on port %s\n", port)
	err := server.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Server crashed: %e", err)
	}

}
