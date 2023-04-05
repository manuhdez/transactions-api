package controller

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/request"
	"github.com/manuhdez/transactions-api/internal/users/infra"
)

type RegisterUser struct {
	service service.RegisterUser
}

func NewRegisterUserController(s service.RegisterUser) RegisterUser {
	return RegisterUser{service: s}
}

func (ct RegisterUser) Handle(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var req request.RegisterUser
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUser := user.New(req.Id, req.FirstName, req.LastName, req.Email, req.Password)
	err = ct.service.Register(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not register user. Please try again."))
		return
	}

	jsonUser := infra.NewUserJson(newUser)
	jsonStr, _ := jsonUser.ToJson()
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(jsonStr)
}
