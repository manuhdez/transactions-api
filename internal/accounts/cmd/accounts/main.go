package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	port := os.Getenv("APP_PORT")
	app := BootstrapApp()

	go app.EventBus.Listen()

	addr := fmt.Sprintf(":%s", port)
	if err := http.ListenAndServe(addr, app.Server.Engine); err != nil {
		log.Fatalf("Accounts service crashed: %s", err)
	}
}
