package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	port := os.Getenv("APP_PORT")
	app := NewApp()

	go app.EventBus.Listen()

	err := app.Server.Engine.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error("Server crashed with ", "error", err.Error())
	}
}
