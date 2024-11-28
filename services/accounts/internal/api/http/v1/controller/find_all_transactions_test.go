package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
)

var userId = "33"

var (
	errFetchingTransactions = errors.New("error fetching transactions")
)

type testSuite struct {
	suite.Suite
	repository *mocks.TransactionRepository
	controller FindAllTransactions
	ctx        *echo.Context
}

func (s *testSuite) SetupTest() {
	s.repository = new(mocks.TransactionRepository)
	retriever := service.NewTransactionsRetriever(s.repository)
	s.controller = NewFindAllTransactions(retriever)
}

func (s *testSuite) TestFindAllSuccess() {
	transactions := []transaction.Transaction{
		transaction.NewDeposit("1", userId, 100),
		transaction.NewWithdraw("1", userId, 125.44),
		transaction.NewDeposit("22", userId, 250),
	}

	s.repository.On("All", mock.Anything, userId).Return(transactions, nil)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, w)
	ctx.Set("userId", userId)

	err := s.controller.Handle(ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusOK, w.Code)

	res, err := json.Marshal(findAllTransactionsResponse{Transactions: transactions})
	require.NoError(s.T(), err)
	assert.JSONEq(s.T(), string(res), w.Body.String())
}

func (s *testSuite) TestFindAllError() {
	s.repository.On("All", mock.Anything, userId).Return(nil, errFetchingTransactions)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	ctx := echo.New().NewContext(req, w)
	ctx.Set("userId", userId)

	err := s.controller.Handle(ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)

}

func TestFindAllTransactionsController(t *testing.T) {
	suite.Run(t, new(testSuite))
}
