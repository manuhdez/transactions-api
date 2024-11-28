package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
)

var (
	errTransfer = errors.New("transfer error")
)

type transferSuite struct {
	bus       *mocks.EventBus
	trxRepo   *mocks.TransactionRepository
	accRepo   *mocks.AccountMockRepository
	transfer  *service.TransactionService
	accFinder *service.AccountFinder
	ctrl      Transfer
}

func setupTransferSuite() transferSuite {
	bus := new(mocks.EventBus)
	trxRepo := new(mocks.TransactionRepository)
	accRepo := new(mocks.AccountMockRepository)
	trx := service.NewTransactionService(trxRepo)
	accFinder := service.NewAccountFinder(accRepo)

	return transferSuite{
		bus:      bus,
		trxRepo:  trxRepo,
		accRepo:  accRepo,
		transfer: trx,
		ctrl:     NewTransferController(trx, accFinder, bus),
	}
}

func (s transferSuite) assertMocks(t *testing.T) {
	s.bus.AssertExpectations(t)
	s.trxRepo.AssertExpectations(t)
}

func TestTransfer_Handle(t *testing.T) {
	e := echo.New()
	e.Validator = sharedhttp.NewRequestValidator()

	originAccount := account.NewWithUserID("1", "999", 0, "EUR")
	destinationAccount := account.NewWithUserID("2", "21", 0, "EUR")

	t.Run("happy path - transfer successful", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		s.accRepo.On("Find", mock.Anything, originAccount.Id()).Return(originAccount, nil).Once()
		s.accRepo.On("Find", mock.Anything, destinationAccount.Id()).Return(destinationAccount, nil).Once()

		withdraw := transaction.NewWithdraw("1", "999", 100)
		s.trxRepo.On("Withdraw", mock.Anything, withdraw).Return(nil)
		deposit := transaction.NewDeposit("2", "999", 100)
		s.trxRepo.On("Deposit", mock.Anything, deposit).Return(nil)
		s.bus.On("Publish", mock.Anything, event.NewWithdrawCreated(withdraw)).Return(nil).Once()
		s.bus.On("Publish", mock.Anything, event.NewDepositCreated(deposit)).Return(nil).Once()

		body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 100})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx := e.NewContext(req, w)
		ctx.Set("userId", "999")

		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("return 400 if request is invalid - missing transfer amount", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		body, err := json.Marshal(transferRequest{From: "1", To: "2"})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		cxt := e.NewContext(req, w)
		err = s.ctrl.Handle(cxt)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("return 401 if user does not own the origin account", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		s.accRepo.On("Find", mock.Anything, mock.Anything).Return(destinationAccount, nil)

		body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 50})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		ctx := e.NewContext(req, w)
		ctx.Set("userId", "999")

		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	t.Run("return 404 if destination account is not found", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		s.accRepo.On("Find", mock.Anything, originAccount.Id()).Return(originAccount, nil).Once()
		s.accRepo.On("Find", mock.Anything, destinationAccount.Id()).Return(account.Account{}, errors.New("not found")).Once()

		body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 50})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		ctx := e.NewContext(req, w)
		ctx.Set("userId", "999")

		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("return 500 if transfer service fails", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		s.accRepo.On("Find", mock.Anything, originAccount.Id()).Return(originAccount, nil).Once()
		s.accRepo.On("Find", mock.Anything, destinationAccount.Id()).Return(destinationAccount, nil).Once()

		s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(errTransfer).Once()
		body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 50})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		ctx := e.NewContext(req, w)
		ctx.Set("userId", "999")

		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
