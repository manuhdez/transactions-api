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
	errTransfer = errors.New("transfer error")
)

type transferSuite struct {
	bus      *mocks.EventBus
	accRepo  *mocks.AccountMockRepository
	trxRepo  *mocks.TransactionMockRepository
	transfer *service.TransactionService
	ctrl     Transfer
}

func setupTransferSuite() transferSuite {
	bus := new(mocks.EventBus)
	accRepo := new(mocks.AccountMockRepository)
	trxRepo := new(mocks.TransactionMockRepository)
	trx := service.NewTransactionService(trxRepo, accRepo)

	return transferSuite{
		bus:      bus,
		accRepo:  accRepo,
		trxRepo:  trxRepo,
		transfer: trx,
		ctrl:     NewTransferController(trx, bus),
	}
}

func (s transferSuite) assertMocks(t *testing.T) {
	s.bus.AssertExpectations(t)
	s.accRepo.AssertExpectations(t)
	s.trxRepo.AssertExpectations(t)
}

func TestTransfer_Handle(t *testing.T) {
	e := echo.New()
	e.Validator = sharedhttp.NewRequestValidator()

	t.Run("happy path - transfer successful", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{Id: "1", UserId: "999"}, nil)
		withdraw := transaction.NewWithdraw("1", "999", 100)
		s.trxRepo.On("Withdraw", mock.Anything, withdraw).Return(nil)
		deposit := transaction.NewDeposit("2", "999", 100)
		s.trxRepo.On("Deposit", mock.Anything, deposit).Return(nil)
		s.bus.On("Publish", mock.Anything, event.NewWithdrawCreated(withdraw)).Return(nil).Once()
		s.bus.On("Publish", mock.Anything, event.NewDepositCreated(deposit)).Return(nil).Once()

		body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 100, UserId: "999"})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewReader(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		ctx := e.NewContext(req, w)

		err = s.ctrl.Handle(ctx)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("return 400 if request is invalid - missing user id", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 50})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		cxt := e.NewContext(req, w)
		err = s.ctrl.Handle(cxt)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("return 500 if transfer service fails", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{}, errTransfer)

		body, err := json.Marshal(transferRequest{UserId: "999", From: "1", To: "2", Amount: 50})
		require.NoError(t, err)

		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		cxt := e.NewContext(req, w)
		err = s.ctrl.Handle(cxt)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})
}
