package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"
)

type Suite struct {
	suite.Suite
	bus        *mocks.EventBus
	accRepo    *mocks.AccountMockRepository
	trxRepo    *mocks.TransactionRepository
	service    *service.TransactionService
	accFinder  *service.AccountFinder
	controller Deposit
	recorder   *httptest.ResponseRecorder
}

func (s *Suite) SetupTest() {
	s.accRepo = new(mocks.AccountMockRepository)
	s.trxRepo = new(mocks.TransactionRepository)
	s.service = service.NewTransactionService(s.trxRepo)
	s.accFinder = service.NewAccountFinder(s.accRepo)
	s.bus = new(mocks.EventBus)

	s.controller = NewDeposit(s.service, s.accFinder, s.bus)
	s.recorder = httptest.NewRecorder()
}

func (s *Suite) assertMocks() {
	s.accRepo.AssertExpectations(s.T())
	s.trxRepo.AssertExpectations(s.T())
	s.bus.AssertExpectations(s.T())
}

func (s *Suite) TestDepositController_Success() {
	userAccount := account.NewWithUserID("1", "999", 0, "EUR")
	deposit := transaction.NewDeposit("1", "999", 100)

	body, err := json.Marshal(depositRequest{Account: deposit.AccountId, Amount: deposit.Amount, Currency: "EUR"})
	if err != nil {
		s.Fail("Error marshaling json")
	}

	s.accRepo.On("Find", mock.Anything, mock.Anything).Return(userAccount, nil).Once()
	s.trxRepo.On("Deposit", mock.Anything, deposit).Return(nil).Once()
	s.bus.On("Publish", mock.Anything, event.NewDepositCreated(deposit)).Return(nil).Once()

	req := httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := echo.New().NewContext(req, s.recorder)
	ctx.Set("userId", "999")
	err = s.controller.Handle(ctx)
	assert.NoError(s.T(), err)

	if s.recorder.Code != 201 {
		s.T().Errorf("Expected status code 201, got %d", s.recorder.Code)
	}
	assert.JSONEq(s.T(), `{"message":"Deposit successfully created"}`, s.recorder.Body.String())
	s.assertMocks()
}

func (s *Suite) TestAccountNotFound() {
	deposit := transaction.NewDeposit("1", "999", 100)

	body, err := json.Marshal(request.Deposit{Account: deposit.AccountId, Amount: deposit.Amount, Currency: "EUR"})
	require.NoError(s.T(), err)

	s.accRepo.On("Find", mock.Anything, mock.Anything).Return(account.Account{}, errAccountNotFound).Once()

	req := httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	ctx := echo.New().NewContext(req, s.recorder)
	ctx.Set("userId", "999")

	err = s.controller.Handle(ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusUnauthorized, s.recorder.Code)
	assert.JSONEq(s.T(), `{"error":"unauthorized"}`, s.recorder.Body.String())
	s.assertMocks()
}

func (s *Suite) TestDepositController_MissingAccount() {
	body, err := json.Marshal(request.Deposit{Amount: 32, Currency: "EUR"})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))
	ctx := echo.New().NewContext(req, s.recorder)
	err = s.controller.Handle(ctx)
	assert.NoError(s.T(), err)

	if s.recorder.Code != 400 {
		s.T().Errorf("Expected status code 400, got %d", s.recorder.Code)
	}
	s.assertMocks()
}

func (s *Suite) TestDepositController_MissingAmount() {
	body, err := json.Marshal(request.Deposit{Account: "112", Currency: "EUR"})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))
	ctx := echo.New().NewContext(req, s.recorder)
	err = s.controller.Handle(ctx)
	assert.NoError(s.T(), err)

	if s.recorder.Code != 400 {
		s.T().Errorf("Expected status code 400, got %d", s.recorder.Code)
	}
	s.assertMocks()
}

func (s *Suite) TestDepositController_MissingCurrency() {
	body, err := json.Marshal(request.Deposit{Account: "123", Amount: 32})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))
	ctx := echo.New().NewContext(req, s.recorder)
	err = s.controller.Handle(ctx)
	assert.NoError(s.T(), err)

	if s.recorder.Code != 400 {
		s.T().Errorf("Expected status code 400, got %d", s.recorder.Code)
	}
	s.assertMocks()
}

func TestDepositController(t *testing.T) {
	suite.Run(t, new(Suite))
}
