package controllers

import (
	"encoding/json"
	"net/http"
)

type statusResponse struct {
	Status string `json:"status"`
}

func StatusController(w http.ResponseWriter, _ *http.Request) {
	r := statusResponse{"ok"}
	response, _ := json.Marshal(r)
	_, _ = w.Write(response)
}
