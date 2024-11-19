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

	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
)

type Suite struct {
	suite.Suite
	service    *mocks.Transactioner
	bus        *mocks.EventBus
	controller controller.Deposit
	recorder   *httptest.ResponseRecorder
}

func (s *Suite) SetupTest() {
	s.service = new(mocks.Transactioner)
	s.bus = new(mocks.EventBus)

	s.controller = controller.NewDeposit(s.service, s.bus)
	s.recorder = httptest.NewRecorder()
}

func (s *Suite) assertMocks() {
	s.service.AssertExpectations(s.T())
	s.bus.AssertExpectations(s.T())
}

func (s *Suite) TestDepositController_Success() {
	trx := transaction.NewDeposit("1", "999", 100)

	body, err := json.Marshal(request.Deposit{Account: trx.AccountId, Amount: trx.Amount, Currency: "EUR"})
	if err != nil {
		s.Fail("Error marshaling json")
	}

	s.service.On("Deposit", mock.Anything, mock.Anything).Return(nil).Once()
	s.service.On("PullEvents").Return([]event.Event{event.NewDepositCreated(trx)})
	s.bus.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()

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
