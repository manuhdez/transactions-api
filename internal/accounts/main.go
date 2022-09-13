package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	server := InitServer()

	go server.EventBus.Listen()

	fmt.Printf("Transactions service running on port %s\n", port)
	err := server.Engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Server crashed: %e", err)
	}
}
