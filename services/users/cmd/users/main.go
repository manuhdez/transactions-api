package main

import (
	"log"
)

func main() {
	app := Init()

	go app.EventBus.Listen()

	if err := app.Server.Listen(); err != nil {
		log.Fatalf("users service crashed: %e", err)
	}
}
