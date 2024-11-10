package controller

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"

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
}

func (ctrl Login) Handle(c echo.Context) error {
	var req request.Login
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	errors := req.Validate()
	if len(errors) > 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"errors": errors})
	}

	user, err := ctrl.loginService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	token, err := ctrl.tokenService.CreateToken(user.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	res, _ := json.Marshal(LoginResponse{
		Success: true,
		UserId:  user.Id,
		Token:   token,
	})
	return c.JSON(http.StatusOK, echo.Map{"data": string(res)})
}
