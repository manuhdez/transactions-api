package controller

import (
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

	return c.JSON(http.StatusOK, echo.Map{"users": getResponseData(users)})
}

func getResponseData(users []user.User) []infra.UserJson {
	res := make([]infra.UserJson, len(users))
	for i := range users {
		res[i] = infra.NewUserJson(users[i])
	}
	return res
}
