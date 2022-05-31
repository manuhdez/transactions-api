package controllers

import (
	"encoding/json"
	"net/http"
)

type StatusController struct{}

func NewStatusController() StatusController {
	return StatusController{}
}

type statusResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func (StatusController) Handle(w http.ResponseWriter, _ *http.Request) {
	status := statusResponse{"ok", "transactions"}
	res, _ := json.Marshal(status)
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
