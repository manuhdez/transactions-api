package controller

import (
	"encoding/json"
	"net/http"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/request"
	"github.com/manuhdez/transactions-api/internal/users/infra"
)

type Login struct {
	loginService service.LoginService
	tokenService infra.TokenService
}

func NewLoginController(login service.LoginService, token infra.TokenService) Login {
	return Login{
		loginService: login,
		tokenService: token,
	}
}

type LoginResponse struct {
	Success bool   `json:"success"`
	UserId  string `json:"id"`
	Token   string `json:"token"`
	// Expiration string `json:"expiration"`
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

	user, err := ctlr.loginService.Login(req.Email, req.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	token, err := ctlr.tokenService.CreateToken(user.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("There was an error creating your session please try again."))
		return
	}

	res, _ := json.Marshal(LoginResponse{
		Success: true,
		UserId:  user.Id,
		Token:   token,
	})
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(res)
}
