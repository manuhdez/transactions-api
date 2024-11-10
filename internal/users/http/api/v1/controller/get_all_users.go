package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

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

func (ctrl GetAllUsers) Handle(c echo.Context) error {
	users, err := ctrl.retriever.Retrieve()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := getResponseData(users)
	response, err := json.Marshal(data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, string(response))
}

func getResponseData(users []user.User) []infra.UserJson {
	var res []infra.UserJson
	for _, u := range users {
		res = append(res, infra.NewUserJson(u))
	}
	return res
}
