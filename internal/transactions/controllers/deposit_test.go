package controllers

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
	"github.com/stretchr/testify/mock"
)

func TestDepositController_Handle(t *testing.T) {

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

	t.Run("Returns a 400 status code if the request is invalid", func(t *testing.T) {
		// Arrange
		repo := new(mocks.TransactionMockRepository)
		ser := service.NewDepositService(repo)
		ctrl := NewDepositController(ser)
		repo.On("Deposit", mock.Anything, mock.Anything).Return(nil)
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		req := bytes.NewBufferString(`{"amount": 100, "currency": "EUR"}`)
		ctx.Request = httptest.NewRequest("POST", "/deposit", req)

		// Act
		ctrl.Handle(ctx)

		// Assert
		if w.Code != 400 {
			t.Errorf("Expected status code 400, got %d", w.Code)
		}
	})
}
