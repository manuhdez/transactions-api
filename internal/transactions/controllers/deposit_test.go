package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
	"github.com/stretchr/testify/mock"
)

func TestDepositController_Success(t *testing.T) {
	t.Run("Returns a 201 status code", func(t *testing.T) {
		// Arrange
		repo := new(mocks.TransactionMockRepository)
		ser := service.NewDepositService(repo)
		ctrl := NewDepositController(ser)
		repo.On("Deposit", mock.Anything, mock.Anything).Return(nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		req := bytes.NewBufferString(`{"account": "123", "amount": 100, "currency": "EUR"}`)
		ctx.Request = httptest.NewRequest("POST", "/deposit", req)

		// Act
		ctrl.Handle(ctx)

		// Assert
		if w.Code != 201 {
			t.Errorf("Expected status code 201, got %d", w.Code)
		}
		if w.Body.String() != `{"message":"Deposit successfully created"}` {
			t.Errorf("Expected body message, got %s", w.Body.String())
		}
	})
}

func TestDepositController_Error(t *testing.T) {
	repo := new(mocks.TransactionMockRepository)
	ser := service.NewDepositService(repo)
	ctrl := NewDepositController(ser)
	repo.On("Deposit", mock.Anything, mock.Anything).Return(nil)

	t.Run("Returns a 400 status if the account id is missing", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)

		body, err := json.Marshal(DepositRequest{Amount: 32, Currency: "EUR"})
		if err != nil {
			t.Fatalf("Error marshaling json: %v", err)
		}
		ctx.Request = httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))

		ctrl.Handle(ctx)

		if w.Code != 400 {
			t.Errorf("Expected status code 400, got %d", w.Code)
		}
	})

	t.Run("Returns a 400 status if the amount is missing", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		body, err := json.Marshal(DepositRequest{Amount: 122, Currency: "EUR"})
		if err != nil {
			t.Errorf("Error marshalling json: %v", err)
		}
		ctx.Request = httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))

		ctrl.Handle(ctx)

		if w.Code != 400 {
			t.Errorf("Expected status code 400, got %d", w.Code)
		}
	})

	t.Run("Returns a 400 status if the currency is missing", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		body, err := json.Marshal(DepositRequest{Account: "123", Amount: 12.3})
		if err != nil {
			t.Fatalf("Error marshalling json: %v", err)
		}
		ctx.Request = httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))

		ctrl.Handle(ctx)

		if w.Code != 400 {
			t.Errorf("Expected status code 400, got %d", w.Code)
		}
	})
}
