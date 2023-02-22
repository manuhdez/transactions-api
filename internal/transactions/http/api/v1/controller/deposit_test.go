package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/controller"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite
	repository *mocks.TransactionMockRepository
	bus        *mocks.EventBus
	controller controller.Deposit
	ctx        *gin.Context
	recorder   *httptest.ResponseRecorder
}

func (s *Suite) SetupTest() {
	s.repository = new(mocks.TransactionMockRepository)
	s.repository.On("Deposit", mock.Anything, mock.Anything).Return(nil)

	s.bus = new(mocks.EventBus)
	s.bus.On("Publish", mock.Anything, mock.Anything).Return(nil)

	s.controller = controller.NewDeposit(service.NewDepositService(s.repository, s.bus))
	s.recorder = httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(s.recorder)
	s.ctx = ctx
}

func (s *Suite) TestDepositController_Success() {
	body, err := json.Marshal(request.Deposit{Account: "333", Amount: 100, Currency: "EUR"})
	if err != nil {
		s.Fail("Error marshaling json")
	}

	s.ctx.Request = httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))
	s.controller.Handle(s.ctx)

	if s.recorder.Code != 201 {
		s.T().Errorf("Expected status code 201, got %d", s.recorder.Code)
	}
	if s.recorder.Body.String() != `{"message":"Deposit successfully created"}` {
		s.T().Errorf("Expected body message, got %s", s.recorder.Body.String())
	}
}

func (s *Suite) TestDepositController_MissingAccount() {
	body, err := json.Marshal(request.Deposit{Amount: 32, Currency: "EUR"})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}
	s.ctx.Request = httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))

	s.controller.Handle(s.ctx)

	if s.recorder.Code != 400 {
		s.T().Errorf("Expected status code 400, got %d", s.recorder.Code)
	}
}

func (s *Suite) TestDepositController_MissingAmount() {
	body, err := json.Marshal(request.Deposit{Account: "112", Currency: "EUR"})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}
	s.ctx.Request = httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))

	s.controller.Handle(s.ctx)

	if s.recorder.Code != 400 {
		s.T().Errorf("Expected status code 400, got %d", s.recorder.Code)
	}
}

func (s *Suite) TestDepositController_MissingCurrency() {
	body, err := json.Marshal(request.Deposit{Account: "123", Amount: 32})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}
	s.ctx.Request = httptest.NewRequest(http.MethodPost, "/deposit", bytes.NewBuffer(body))

	s.controller.Handle(s.ctx)

	if s.recorder.Code != 400 {
		s.T().Errorf("Expected status code 400, got %d", s.recorder.Code)
	}
}

func TestDepositController(t *testing.T) {
	suite.Run(t, new(Suite))
}
