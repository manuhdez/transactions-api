package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

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

func (ctrl RegisterUser) Handle(c echo.Context) error {
	var req request.RegisterUser
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	errors := req.Validate()
	if errs := len(errors); errs > 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": req.ErrorResponse(errors)})
	}

	newUser := user.New(req.Id, req.FirstName, req.LastName, req.Email, req.Password)
	if err := ctrl.service.Register(newUser); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, infra.NewUserJson(newUser))
}
