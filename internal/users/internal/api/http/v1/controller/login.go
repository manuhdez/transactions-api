package controller

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/users/internal/api/http/v1/request"
	"github.com/manuhdez/transactions-api/internal/users/internal/application/service"
	domain_service "github.com/manuhdez/transactions-api/internal/users/internal/domain/service"
)

var (
	errInvalidLoginRequest = "invalid request body"
	errMissingCredentials  = "missing required email or password field"
	errInvalidCredentials  = "invalid credentials"
	errSessionCreate       = "there was an issue creating your session"
)

type Login struct {
	loginService service.LoginService
	tokenService domain_service.TokenService
}

func NewLoginController(login service.LoginService, token domain_service.TokenService) Login {
	return Login{
		loginService: login,
		tokenService: token,
	}
}

type LoginResponse struct {
	Success bool   `json:"success"`
	UserId  string `json:"id,omitempty" `
	Token   string `json:"token,omitempty"`
	Error   string `json:"error,omitempty"`
}

func (ctrl Login) Handle(c echo.Context) error {
	var req request.Login
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, LoginResponse{Success: false, Error: errInvalidLoginRequest})
	}

	if err := c.Validate(&req); err != nil {
		log.Printf("[Login:Handle][err: %s]", err)
		return c.JSON(http.StatusBadRequest, LoginResponse{Success: false, Error: errMissingCredentials})
	}

	user, err := ctrl.loginService.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, LoginResponse{Success: false, Error: errInvalidCredentials})
	}

	token, err := ctrl.tokenService.CreateToken(user.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, LoginResponse{Success: false, Error: errSessionCreate})
	}

	return c.JSON(http.StatusOK, LoginResponse{
		Success: true,
		UserId:  user.Id,
		Token:   token,
	})
}
