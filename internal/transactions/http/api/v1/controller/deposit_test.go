package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
)

type Suite struct {
	suite.Suite
	repository *mocks.TransactionMockRepository
	accRepo    *mocks.AccountMockRepository
	bus        *mocks.EventBus
	controller controller.Deposit
	recorder   *httptest.ResponseRecorder
}

func (s *Suite) SetupTest() {
	s.repository = new(mocks.TransactionMockRepository)
	s.repository.On("Deposit", mock.Anything, mock.Anything).Return(nil)

	s.accRepo = new(mocks.AccountMockRepository)
	s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{UserId: "999"}, nil)

	s.bus = new(mocks.EventBus)
	s.bus.On("Publish", mock.Anything, mock.Anything).Return(nil)

	s.controller = controller.NewDeposit(service.NewTransactionService(s.repository, s.accRepo, s.bus))
	s.recorder = httptest.NewRecorder()
}

func (s *Suite) TestDepositController_Success() {
	body, err := json.Marshal(request.Deposit{Account: "333", Amount: 100, Currency: "EUR"})
	if err != nil {
		s.Fail("Error marshaling json")
	}

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
}

func TestDepositController(t *testing.T) {
	suite.Run(t, new(Suite))
}
