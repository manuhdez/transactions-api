package controller

import (
	"encoding/json"
	"net/http"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/request"
)

type Login struct {
	service service.LoginService
}

func NewLoginController(s service.LoginService) Login {
	return Login{s}
}

type LoginResponse struct {
	Success    bool   `json:"success"`
	UserId     string `json:"id"`
	Token      string `json:"token"`
	Expiration string `json:"expiration"`
}

func (ctlr Login) Handle(w http.ResponseWriter, r *http.Request) {
	var req request.Login
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.Validate()
	if len(req.Errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(req.ErrorResponse()))
		return
	}

	user, err := ctlr.service.Login(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	response := LoginResponse{true, user.Id, "", ""}
	res, err := json.Marshal(response)

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
