package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/users/internal/application/dtos"
	"github.com/manuhdez/transactions-api/internal/users/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/users/internal/domain/user"
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

func getResponseData(users []user.User) []dtos.UserJson {
	res := make([]dtos.UserJson, len(users))
	for i := range users {
		res[i] = dtos.NewUserJson(users[i])
	}
	return res
}
