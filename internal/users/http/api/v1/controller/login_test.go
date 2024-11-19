package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/request"
	"github.com/manuhdez/transactions-api/internal/users/test/mocks"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
)

type loginSuite struct {
	userRepo *mocks.UserRepository
	hasher   *mocks.Hasher
	token    *mocks.TokenService
	login    service.LoginService
	ctrl     Login
}

func setupLoginSuite() loginSuite {
	userRepo := new(mocks.UserRepository)
	hasher := new(mocks.Hasher)
	tokenSrv := new(mocks.TokenService)
	loginSrv := service.NewLoginService(userRepo, hasher)

	return loginSuite{
		userRepo: userRepo,
		hasher:   hasher,
		token:    tokenSrv,
		login:    loginSrv,
		ctrl:     NewLoginController(loginSrv, tokenSrv),
	}
}

func (s loginSuite) assertMocks(t *testing.T) {
	s.userRepo.AssertExpectations(t)
	s.hasher.AssertExpectations(t)
	s.token.AssertExpectations(t)
}

func TestLogin_Handle(t *testing.T) {
	e := echo.New()
	e.Validator = sharedhttp.NewRequestValidator()

	t.Run("Happy path - login with valid credentials", func(t *testing.T) {
		s := setupLoginSuite()
		defer s.assertMocks(t)

		s.userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(user.User{Id: "1", Email: "test@mail.com"}, nil)
		s.hasher.On("Compare", mock.Anything, mock.Anything).Return(nil)
		s.token.On("CreateToken", mock.Anything).Return("mock-token", nil)

		body, err := json.Marshal(request.Login{Email: "test@mail.com", Password: "test-password"})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		ctx := e.NewContext(req, w)
		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)

		res, err := json.Marshal(LoginResponse{Success: true, UserId: "1", Token: "mock-token"})
		require.NoError(t, err)
		assert.JSONEq(t, string(res), w.Body.String())
	})

	t.Run("login fail - invalid request", func(t *testing.T) {
		s := setupLoginSuite()
		defer s.assertMocks(t)

		body, err := json.Marshal(request.Login{Email: "test@mail.com"})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		ctx := e.NewContext(req, w)
		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)

		res, err := json.Marshal(LoginResponse{Success: false, Error: errMissingCredentials})
		require.NoError(t, err)
		assert.JSONEq(t, string(res), w.Body.String())
	})

	t.Run("login fail - invalid email (user not found)", func(t *testing.T) {
		s := setupLoginSuite()
		defer s.assertMocks(t)

		s.userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(user.User{}, errors.New("user not found"))

		body, err := json.Marshal(request.Login{Email: "test@mail.com", Password: "test-password"})
		require.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		w := httptest.NewRecorder()
		ctx := e.NewContext(req, w)

		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, w.Code)

		res, err := json.Marshal(LoginResponse{Success: false, Error: errInvalidCredentials})
		require.NoError(t, err)
		assert.JSONEq(t, string(res), w.Body.String())
	})

	t.Run("login fail - wrong password", func(t *testing.T) {
		s := setupLoginSuite()
		defer s.assertMocks(t)

		s.userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(user.User{Id: "1", Email: "test@email.com"}, nil)
		s.hasher.On("Compare", mock.Anything, mock.Anything).Return(errors.New("invalid password"))

		body, err := json.Marshal(request.Login{Email: "test@email.com", Password: "password"})
		require.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()

		ctx := e.NewContext(req, w)
		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, w.Code)

		res, err := json.Marshal(LoginResponse{Success: false, Error: errInvalidCredentials})
		require.NoError(t, err)
		assert.JSONEq(t, string(res), w.Body.String())
	})

	t.Run("login fail - error generating token", func(t *testing.T) {
		s := setupLoginSuite()
		defer s.assertMocks(t)

		s.userRepo.On("FindByEmail", mock.Anything, mock.Anything).Return(user.User{Id: "1", Email: "test@email.com"}, nil)
		s.hasher.On("Compare", mock.Anything, mock.Anything).Return(nil)
		s.token.On("CreateToken", mock.Anything).Return("", errors.New("error generating token"))

		body, err := json.Marshal(request.Login{Email: "test@email.com", Password: "password"})
		require.NoError(t, err)
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()

		ctx := e.NewContext(req, w)
		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)

		res, err := json.Marshal(LoginResponse{Success: false, Error: errSessionCreate})
		require.NoError(t, err)
		assert.JSONEq(t, string(res), w.Body.String())
	})
}
