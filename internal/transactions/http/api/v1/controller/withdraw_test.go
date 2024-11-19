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
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
)

type withDrawSuite struct {
	suite.Suite

	service    *mocks.Transactioner
	bus        *mocks.EventBus
	controller controller.Withdraw
	recorder   *httptest.ResponseRecorder
	server     *echo.Echo
}

func (s *withDrawSuite) SetupTest() {
	s.service = new(mocks.Transactioner)
	s.bus = new(mocks.EventBus)
	s.controller = controller.NewWithdraw(s.service, s.bus)

	e := echo.New()
	e.Validator = sharedhttp.NewRequestValidator()
	s.server = e
	s.recorder = httptest.NewRecorder()
}

func (s *withDrawSuite) assertMocks() {
	s.service.AssertExpectations(s.T())
	s.bus.AssertExpectations(s.T())
}

func (s *withDrawSuite) TestWithdrawController_Success() {
	trx := transaction.NewWithdraw("1", "999", 125)
	body, err := json.Marshal(request.Withdraw{Account: trx.AccountId, Amount: trx.Amount, Currency: "EUR"})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	s.service.On("Withdraw", mock.Anything, trx).Return(nil).Once()
	s.service.On("PullEvents").Return([]event.Event{event.NewWithdrawCreated(trx)}).Once()
	s.bus.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()

	ctx := s.server.NewContext(req, s.recorder)
	ctx.Set("userId", "999")
	err = s.controller.Handle(ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 201, s.recorder.Code)
	s.assertMocks()
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
	s.assertMocks()
}

func TestWithdrawController(t *testing.T) {
	suite.Run(t, new(withDrawSuite))
}
