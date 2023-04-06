package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	app := NewApp()

	go app.EventBus.Listen()

	fmt.Printf("Transactions service running on port %s\n", port)
	err := app.Server.Engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Server crashed: %e", err)
	}
}
