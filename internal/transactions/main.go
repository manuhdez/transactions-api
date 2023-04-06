package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	server := NewServer()

	go server.EventBus.Listen()

	err := server.Engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("transactions service crashed: %s", err.Error())
	}
}
