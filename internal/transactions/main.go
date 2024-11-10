package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	app := NewApp()

	go app.EventBus.Listen()

	addr := fmt.Sprintf(":%s", port)
	if err := http.ListenAndServe(addr, app.Server.Engine); err != nil {
		log.Fatalf("transactions service crashed: %s", err.Error())
	}
}
