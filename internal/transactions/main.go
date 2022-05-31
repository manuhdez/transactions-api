package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	server := InitServer()
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), server)
	if err != nil {
		panic(err)
	}
}
