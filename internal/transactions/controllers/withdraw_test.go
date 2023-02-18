package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
)

type withDrawSuite struct {
	suite.Suite
	controller WithdrawController
	ctx        *gin.Context
	recorder   *httptest.ResponseRecorder
}

func (s *withDrawSuite) SetupTest() {
	repository := new(mocks.TransactionMockRepository)
	repository.On("Withdraw", mock.Anything, mock.Anything).Return(nil)

	bus := new(mocks.EventBus)
	bus.On("Publish", mock.Anything, mock.Anything).Return(nil)

	srv := service.NewWithdrawService(repository, bus)
	s.controller = NewWithdrawController(srv)
	s.recorder = httptest.NewRecorder()

	ctx, _ := gin.CreateTestContext(s.recorder)
	s.ctx = ctx
}

func (s *withDrawSuite) TestWithdrawController_Success() {
	body, err := json.Marshal(WithdrawRequest{Account: "112", Amount: 125, Currency: "EUR"})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	s.ctx.Request = httptest.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(body))
	s.controller.Handle(s.ctx)

	assert.Equal(s.T(), 201, s.recorder.Code)
}

func (s *withDrawSuite) TestWithdrawController_BadRequest() {
	body, err := json.Marshal(WithdrawRequest{Account: "112", Amount: 125})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	s.ctx.Request = httptest.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(body))
	s.controller.Handle(s.ctx)
	assert.Equal(s.T(), 400, s.recorder.Code)
}

func TestWithdrawController(t *testing.T) {
	suite.Run(t, new(withDrawSuite))
}
