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
	transfer *mocks.ServiceTransferer
	ctrl     Transfer
}

func setupTransferSuite() transferSuite {
	bus := new(mocks.EventBus)
	trx := new(mocks.ServiceTransferer)

	return transferSuite{
		bus:      bus,
		transfer: trx,
		ctrl:     NewTransferController(trx, bus),
	}
}

func (s transferSuite) assertMocks(t *testing.T) {
	s.bus.AssertExpectations(t)
	s.transfer.AssertExpectations(t)
}

func TestTransfer_Handle(t *testing.T) {
	e := echo.New()
	e.Validator = sharedhttp.NewRequestValidator()

	t.Run("happy path - transfer successful", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		s.transfer.On("Transfer", mock.Anything, mock.Anything).Return(nil).Once()
		s.transfer.On("PullEvents").Return([]event.Event{
			event.NewWithdrawCreated(transaction.NewWithdraw("1", "999", 100)),
			event.NewDepositCreated(transaction.NewDeposit("2", "999", 100)),
		}).Once()
		s.bus.On("Publish", mock.Anything, mock.Anything).Return(nil).Twice()

		body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 50, UserId: "999"})
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

		s.transfer.On("Transfer", mock.Anything, mock.Anything).Return(errTransfer).Once()

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
