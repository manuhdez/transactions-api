package controller_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/request"
	"github.com/manuhdez/transactions-api/internal/users/infra"
	"github.com/manuhdez/transactions-api/internal/users/test/mocks"
)

func TestRegisterUser_Handle(t *testing.T) {
	repo := mocks.UserMockRepository{Err: nil}
	hasher := infra.NewBcryptService()
	bus := new(mocks.Bus)
	bus.On("Publish", mock.Anything, mock.Anything).Return(nil)
	serv := service.NewRegisterUserService(repo, hasher, bus)
	ctrl := controller.NewRegisterUserController(serv)

	t.Run("returns status 201", func(t *testing.T) {
		bus.On("Publish", mock.Anything, mock.Anything).Return(nil)
		reqData := getValidRequest(t)
		req := httptest.NewRequest(http.MethodPost, "/", reqData)
		recorder := httptest.NewRecorder()
		ctrl.Handle(recorder, req)

		got := recorder.Result().StatusCode
		want := http.StatusCreated
		if got != want {
			t.Errorf("Http Status: got %d want %d", got, want)
		}
	})

	t.Run("returns status 400 with no body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		recorder := httptest.NewRecorder()
		ctrl.Handle(recorder, req)

		got := recorder.Result().StatusCode
		want := http.StatusBadRequest
		if got != want {
			t.Errorf("Http Status: got %d want %d", got, want)
		}
	})

	t.Run("returns status 400 with missing fields", func(t *testing.T) {
		reqData := getMissingFieldRequest(t)
		req := httptest.NewRequest(http.MethodPost, "/", reqData)
		recorder := httptest.NewRecorder()
		ctrl.Handle(recorder, req)

		got := recorder.Result().StatusCode
		want := http.StatusBadRequest
		if got != want {
			t.Errorf("Http Status: got %d want %d", got, want)
		}
	})

	t.Run("returns the newly created user json", func(t *testing.T) {
		reqData := getValidRequest(t)
		req := httptest.NewRequest(http.MethodPost, "/", reqData)
		recorder := httptest.NewRecorder()
		ctrl.Handle(recorder, req)

		body, _ := io.ReadAll(recorder.Body)

		got := string(body)
		want := `{"id":"123-456","first_name":"Ramon","last_name":"Perez","email":"ramon@perez.com"}`
		if got != want {
			t.Errorf("Http Status: got %s want %s", got, want)
		}
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
