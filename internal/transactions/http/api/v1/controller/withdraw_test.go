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
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/http/api/v1/request"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
)

type withDrawSuite struct {
	suite.Suite
	controller Withdraw
	recorder   *httptest.ResponseRecorder
	server     *echo.Echo
}

func (s *withDrawSuite) SetupTest() {
	repository := new(mocks.TransactionMockRepository)
	repository.On("Withdraw", mock.Anything, mock.Anything).Return(nil)

	bus := new(mocks.EventBus)
	bus.On("Publish", mock.Anything, mock.Anything).Return(nil)

	srv := service.NewWithdrawService(repository, bus)
	s.controller = NewWithdraw(srv)
	s.recorder = httptest.NewRecorder()

	e := echo.New()
	e.Validator = sharedhttp.NewRequestValidator()
	s.server = e
}

func (s *withDrawSuite) TestWithdrawController_Success() {
	body, err := json.Marshal(request.Withdraw{Account: "112", Amount: 125, Currency: "EUR"})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	ctx := s.server.NewContext(req, s.recorder)
	err = s.controller.Handle(ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 201, s.recorder.Code)
}

func (s *withDrawSuite) TestWithdrawController_BadRequest() {
	body, err := json.Marshal(request.Withdraw{Account: "112", Amount: 125})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := s.server.NewContext(req, s.recorder)
	err = s.controller.Handle(ctx)
	assert.Equal(s.T(), 400, s.recorder.Code)
}

func TestWithdrawController(t *testing.T) {
	suite.Run(t, new(withDrawSuite))
}
