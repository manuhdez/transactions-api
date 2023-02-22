package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	repository *mocks.TransactionMockRepository
	controller FindAllTransactions
	ctx        *gin.Context
	recorder   *httptest.ResponseRecorder
}

func (s *testSuite) SetupTest() {
	s.repository = new(mocks.TransactionMockRepository)

	s.controller = NewFindAllTransactions(service.NewFindAllTransactionsService(s.repository))
	s.recorder = httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(s.recorder)
	s.ctx = ctx
}

func (s *testSuite) TestFindAllSuccess() {
	expected := http.StatusOK

	s.repository.On("FindAll", mock.Anything, mock.Anything).Return([]transaction.Transaction{}, nil)
	s.ctx.Request = httptest.NewRequest(http.MethodGet, "/transactions", nil)
	s.controller.Handle(s.ctx)

	if s.recorder.Code != expected {
		s.T().Errorf("Expected status code %d, got %d", expected, s.recorder.Code)
	}
}

func (s *testSuite) TestFindAllError() {
	expected := http.StatusInternalServerError
	s.repository.On("FindAll", mock.Anything, mock.Anything).Return([]transaction.Transaction{}, errors.New("there was an error"))
	s.ctx.Request = httptest.NewRequest(http.MethodGet, "/transactions", nil)
	s.controller.Handle(s.ctx)

	if s.recorder.Code != expected {
		s.T().Errorf("Expected status code %d, got %d", expected, s.recorder.Code)
	}
}

func TestFindAllTransactionsController(t *testing.T) {
	suite.Run(t, new(testSuite))
}
