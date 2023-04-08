package controller

import (
	"encoding/json"
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
	var req request.RegisterUser
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	req.Validate()
	if errs := len(req.Errors); errs > 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(req.ErrorResponse()))
		return
	}

	newUser := user.New(req.Id, req.FirstName, req.LastName, req.Email, req.Password)
	err = ct.service.Register(newUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Could not register user. Please try again."))
		return
	} else {
		w.WriteHeader(http.StatusCreated)
	}

	jsonUser, _ := infra.NewUserJson(newUser).ToJson()
	_, _ = w.Write(jsonUser)
}
