package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type createAccountRequest struct {
	Id      string  `json:"id"`
	Balance float32 `json:"balance"`
}

func CreateAccountController(service service.CreateService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var request createAccountRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		acc := account.New(request.Id, request.Balance)
		err = service.Create(acc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
