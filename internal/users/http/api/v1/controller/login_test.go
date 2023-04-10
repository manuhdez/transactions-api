package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

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

type mockTokenService struct{}

func (m mockTokenService) CreateToken(_ string) (string, error) {
	return "mock-token", nil
}

func (m mockTokenService) ValidateToken(_ string) bool {
	return false
}

func TestLogin_Handle(t *testing.T) {
	repo := mocks.UserMockRepository{Err: nil}
	loginService := service.NewLoginService(repo)
	tokenService := mockTokenService{}
	ctrl := controller.NewLoginController(loginService, tokenService)

	for _, tt := range []testCase{
		{
			name:     "with valid credentials",
			request:  request.Login{Email: "test@mail.com", Password: "test-password"},
			status:   200,
			response: controller.LoginResponse{Success: true, UserId: "1", Token: "mock-token"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			body, err := json.Marshal(tt.request)
			if err != nil {
				t.Fatal("Failed to marshal request: ", err)
			}

			req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(body))
			recorder := httptest.NewRecorder()
			ctrl.Handle(recorder, req)

			status := recorder.Result().StatusCode
			if status != tt.status {
				t.Errorf("Login Status: got %d, want %d", status, tt.status)
			}

			res, err := io.ReadAll(recorder.Result().Body)
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
