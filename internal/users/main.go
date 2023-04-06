package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	app := Init()

	fmt.Printf("Users service running on http://localhost:%s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), app.Server.Engine); err != nil {
		log.Fatalf("users service crashed: %e", err)
	}
}
