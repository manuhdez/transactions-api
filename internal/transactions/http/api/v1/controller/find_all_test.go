package controller

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
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

	s.repository.On("FindAll", mock.Anything, mock.Anything).Return(
		[]transaction.Transaction{{Type: transaction.Withdrawal, Amount: 125.44, AccountId: "22"}},
		nil,
	)

	req := httptest.NewRequest(http.MethodGet, "/transactions", nil)
	ctx := echo.New().NewContext(req, s.recorder)
	err := s.controller.Handle(ctx)
	assert.NoError(s.T(), err)

	if s.recorder.Code != expected {
		s.T().Errorf("Expected status code %d, got %d", expected, s.recorder.Code)
	}

	body, err := io.ReadAll(s.recorder.Body)
	if err != nil {
		s.T().Errorf("io.ReadAll(Body): Unable to read the response body. \n %e", err)
	}

	got := strings.Contains(string(body), "Currency")
	assert.Equal(s.T(), got, false, fmt.Sprintf("body: %s \nShould not contain 'currency' field", string(body)))
}

func (s *testSuite) TestFindAllError() {
	expected := http.StatusInternalServerError
	s.repository.On("FindAll", mock.Anything, mock.Anything).Return([]transaction.Transaction{}, errors.New("there was an error"))
	req := httptest.NewRequest(http.MethodGet, "/transactions", nil)
	ctx := echo.New().NewContext(req, s.recorder)
	err := s.controller.Handle(ctx)
	assert.NoError(s.T(), err)

	if s.recorder.Code != expected {
		s.T().Errorf("Expected status code %d, got %d", expected, s.recorder.Code)
	}
}

func TestFindAllTransactionsController(t *testing.T) {
	suite.Run(t, new(testSuite))
}
