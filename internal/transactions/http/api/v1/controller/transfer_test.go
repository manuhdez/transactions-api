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

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
)

var (
	errTransfer        = errors.New("transfer error")
	errAccountNotFound = errors.New("account not found")
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

	originAccount := account.NewAccount("1", "999")
	destinationAccount := account.NewAccount("2", "21")

	t.Run("happy path - transfer successful", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		s.accRepo.On("FindById", mock.Anything, originAccount.Id).Return(originAccount, nil).Once()
		s.accRepo.On("FindById", mock.Anything, destinationAccount.Id).Return(destinationAccount, nil).Once()

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

		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(destinationAccount, nil)

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

		s.accRepo.On("FindById", mock.Anything, originAccount.Id).Return(originAccount, nil).Once()
		s.accRepo.On("FindById", mock.Anything, destinationAccount.Id).Return(account.Account{}, errAccountNotFound).Once()

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

		s.accRepo.On("FindById", mock.Anything, originAccount.Id).Return(originAccount, nil).Once()
		s.accRepo.On("FindById", mock.Anything, destinationAccount.Id).Return(destinationAccount, nil).Once()

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
