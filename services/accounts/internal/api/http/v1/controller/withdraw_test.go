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
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
)

type withDrawSuite struct {
	suite.Suite

	accRepo    *mocks.AccountRepository
	trxRepo    *mocks.TransactionRepository
	service    *service.WithdrawService
	accFinder  *service.AccountFinder
	bus        *mocks.EventBus
	controller Withdraw
	recorder   *httptest.ResponseRecorder
	server     *echo.Echo
}

func (s *withDrawSuite) SetupTest() {
	s.accRepo = new(mocks.AccountRepository)
	s.trxRepo = new(mocks.TransactionRepository)
	s.bus = new(mocks.EventBus)

	s.service = service.NewWithdrawService(s.trxRepo, s.bus)
	s.accFinder = service.NewAccountFinder(s.accRepo)

	s.controller = NewWithdraw(s.service, s.accFinder)

	e := echo.New()
	e.Validator = sharedhttp.NewRequestValidator()
	s.server = e
	s.recorder = httptest.NewRecorder()
}

func (s *withDrawSuite) assertMocks() {
	s.accRepo.AssertExpectations(s.T())
	s.trxRepo.AssertExpectations(s.T())
	s.bus.AssertExpectations(s.T())
}

func (s *withDrawSuite) TestWithdrawController_Success() {
	userAccount := account.NewWithUserID("1", "999", 0, "EUR")
	trx := transaction.NewWithdraw("1", "999", 125)
	body, err := json.Marshal(withdrawRequest{Account: trx.AccountId, Amount: trx.Amount, Currency: "EUR"})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	s.accRepo.On("Find", mock.Anything, mock.Anything).Return(userAccount, nil).Once()
	s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(nil).Once()
	s.bus.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()

	ctx := s.server.NewContext(req, s.recorder)
	ctx.Set("userId", "999")
	err = s.controller.Handle(ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 201, s.recorder.Code)
	s.assertMocks()
}

func (s *withDrawSuite) TestAccountNotFound() {
	trx := transaction.NewWithdraw("1", "999", 125)
	body, err := json.Marshal(withdrawRequest{Account: trx.AccountId, Amount: trx.Amount, Currency: "EUR"})
	require.NoError(s.T(), err)

	req := httptest.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	s.accRepo.On("Find", mock.Anything, mock.Anything).Return(account.Account{}, errAccountNotFound).Once()

	ctx := s.server.NewContext(req, s.recorder)
	ctx.Set("userId", "999")
	err = s.controller.Handle(ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusUnauthorized, s.recorder.Code)
	s.assertMocks()
}

func (s *withDrawSuite) TestWithdrawController_BadRequest() {
	body, err := json.Marshal(withdrawRequest{Account: "112", Amount: 125})
	if err != nil {
		s.T().Fatalf("Error marshaling json: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/withdraw", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := s.server.NewContext(req, s.recorder)
	assert.NoError(s.T(), s.controller.Handle(ctx))
	assert.Equal(s.T(), 400, s.recorder.Code)
	s.assertMocks()
}

func TestWithdrawController(t *testing.T) {
	suite.Run(t, new(withDrawSuite))
}
