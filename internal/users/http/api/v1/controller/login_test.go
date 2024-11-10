package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/request"
	"github.com/manuhdez/transactions-api/internal/users/test/mocks"
)

type testCase struct {
	name     string
	request  request.Login
	status   int
	response controller.LoginResponse
}

func TestLogin_Handle(t *testing.T) {
	repo := mocks.UserMockRepository{Err: nil}
	hasher := mocks.HasherService{HashError: nil}
	loginService := service.NewLoginService(repo, hasher)
	tokenService := mocks.NewTokenService("mock-token")
	ctrl := controller.NewLoginController(loginService, tokenService)

	for _, tt := range []testCase{
		{
			name:     "with valid credentials",
			request:  request.Login{Email: "test@mail.com", Password: "test-password"},
			status:   http.StatusOK,
			response: controller.LoginResponse{Success: true, UserId: "1", Token: "mock-token"},
		},
		{
			name:     "with wrong credentials",
			request:  request.Login{Email: "test@mail.com", Password: "wrong-password"},
			status:   http.StatusBadRequest,
			response: controller.LoginResponse{Success: false},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatal("Failed to marshal request: ", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			ctx := echo.New().NewContext(req, w)
			err = ctrl.Handle(ctx)
			assert.NoError(t, err)

			if status := w.Result().StatusCode; status != tt.status {
				t.Errorf("Login Status: got %d, want %d", status, tt.status)
			}

			res, err := io.ReadAll(w.Result().Body)
			if err != nil {
				t.Errorf("failed to read body response: %s", err)
			}

			var response controller.LoginResponse
			err = json.Unmarshal(res, &response)
			if response != tt.response {
				t.Errorf("Login Response: got %v, want %v", response, tt.response)
			}
		})
	}
}
