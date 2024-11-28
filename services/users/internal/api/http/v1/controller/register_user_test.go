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
	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/users/internal/api/http/v1/controller"
	"github.com/manuhdez/transactions-api/internal/users/internal/api/http/v1/request"
	"github.com/manuhdez/transactions-api/internal/users/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/users/internal/infra"
	"github.com/manuhdez/transactions-api/internal/users/test/mocks"
)

func TestRegisterUser_Handle(t *testing.T) {
	repo := new(mocks.UserRepository)
	hasher := infra.NewBcryptService()
	bus := new(mocks.EventBus)
	bus.On("Publish", mock.Anything, mock.Anything).Return(nil)
	serv := service.NewRegisterUserService(repo, hasher, bus)
	ctrl := controller.NewRegisterUserController(serv)

	t.Run("returns status 201", func(t *testing.T) {
		repo.On("Save", mock.Anything, mock.Anything).Return(nil)
		bus.On("Publish", mock.Anything, mock.Anything).Return(nil)
		reqData := getValidRequest(t)
		req := httptest.NewRequest(http.MethodPost, "/", reqData)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, w)
		err := ctrl.Handle(ctx)
		assert.NoError(t, err)

		got := w.Result().StatusCode
		want := http.StatusCreated
		if got != want {
			t.Errorf("Http Status: got %d want %d", got, want)
		}
	})

	t.Run("returns status 400 with no body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, w)
		err := ctrl.Handle(ctx)
		assert.NoError(t, err)

		got := w.Result().StatusCode
		want := http.StatusBadRequest
		if got != want {
			t.Errorf("Http Status: got %d want %d", got, want)
		}
	})

	t.Run("returns status 400 with missing fields", func(t *testing.T) {
		reqData := getMissingFieldRequest(t)
		req := httptest.NewRequest(http.MethodPost, "/", reqData)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, w)
		err := ctrl.Handle(ctx)
		assert.NoError(t, err)

		got := w.Result().StatusCode
		want := http.StatusBadRequest
		if got != want {
			t.Errorf("Http Status: got %d want %d", got, want)
		}
	})

	t.Run("returns the newly created user json", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", getValidRequest(t))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		w := httptest.NewRecorder()
		ctx := echo.New().NewContext(req, w)
		err := ctrl.Handle(ctx)
		assert.NoError(t, err)

		body, _ := io.ReadAll(w.Body)

		want := `{"id":"123-456","first_name":"Ramon","last_name":"Perez","email":"ramon@perez.com"}`
		assert.JSONEq(t, want, string(body))
	})
}

func getValidRequest(t *testing.T) io.Reader {
	req, err := json.Marshal(request.RegisterUser{
		Id:        "123-456",
		FirstName: "Ramon",
		LastName:  "Perez",
		Email:     "ramon@perez.com",
		Password:  "my_pass",
	})

	if err != nil {
		t.Fatalf("Failed to create signup request")
		return nil
	}

	return bytes.NewBuffer(req)
}

func getMissingFieldRequest(t *testing.T) io.Reader {
	req, err := json.Marshal(request.RegisterUser{
		Id:        "123-098",
		FirstName: "Ramon",
		LastName:  "Perez",
		Password:  "hello_world",
	})

	if err != nil {
		t.Fatalf("Failed to create request")
		return nil
	}

	return bytes.NewBuffer(req)
}
