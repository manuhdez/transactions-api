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

	err := app.Server.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("transactions service crashed: %s", err.Error())
	}
}
