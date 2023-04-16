package controller

import (
	"encoding/json"
	"net/http"
)

type HealthCheck struct{}

func NewHealthCheck() HealthCheck {
	return HealthCheck{}
}

type HealthCheckResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func (c HealthCheck) Handle(w http.ResponseWriter, _ *http.Request) {

	body, _ := json.Marshal(HealthCheckResponse{Status: "ok", Service: "users"})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}
