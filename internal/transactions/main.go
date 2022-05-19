package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("APP_PORT")
	server := http.NewServeMux()
	server.HandleFunc("/status", statusHandler)

	fmt.Printf("Transactions service running on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), server)
	if err != nil {
		panic(err)
	}
}

type statusRes struct {
	Status string `json:"status"`
}

func statusHandler(w http.ResponseWriter, _ *http.Request) {
	status := statusRes{"ok"}
	res, _ := json.Marshal(status)
	_, _ = w.Write(res)
}
