package main

import (
	"log"
)

func main() {
	app := Init()

	if err := app.Server.Listen(); err != nil {
		log.Fatalf("users service crashed: %e", err)
	}
}
