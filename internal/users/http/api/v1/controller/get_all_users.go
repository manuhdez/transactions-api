package controller

import (
	"encoding/json"
	"net/http"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/infra"
)

type GetAllUsers struct {
	retriever service.UsersRetriever
}

func NewGetAllUsersController(retriever service.UsersRetriever) GetAllUsers {
	return GetAllUsers{
		retriever: retriever,
	}
}

func (ctlr GetAllUsers) Handle(w http.ResponseWriter, _ *http.Request) {
	users, err := ctlr.retriever.Retrieve()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if len(users) == 0 {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("[]"))
		return
	}

	data := getResponseData(users)
	response, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(response)
	return
}

func getResponseData(users []user.User) []infra.UserJson {
	var res []infra.UserJson
	for _, u := range users {
		res = append(res, infra.NewUserJson(u))
	}
	return res
}
