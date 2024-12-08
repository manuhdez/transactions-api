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
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"
)

var (
	errTransfer = errors.New("transfer error")
)

type transferSuite struct {
	trxRepo   *mocks.TransactionRepository
	accRepo   *mocks.AccountRepository
	publisher *mocks.EventBus
	ctrl      Transfer
}

func setupTransferSuite() transferSuite {
	trxRepo := new(mocks.TransactionRepository)
	accRepo := new(mocks.AccountRepository)
	publisher := new(mocks.EventBus)

	transferService := service.NewTransferService(trxRepo, publisher)
	accFinder := service.NewAccountFinder(accRepo)

	return transferSuite{
		trxRepo:   trxRepo,
		accRepo:   accRepo,
		publisher: publisher,
		ctrl:      NewTransferController(transferService, accFinder),
	}
}

func (s transferSuite) assertMocks(t *testing.T) {
	s.trxRepo.AssertExpectations(t)
	s.accRepo.AssertExpectations(t)
	s.publisher.AssertExpectations(t)
}

func TestTransfer_Handle(t *testing.T) {
	e := echo.New()
	e.Validator = sharedhttp.NewRequestValidator()

	originAccount := account.NewWithUserID("1", "999", 500, "EUR")
	destinationAccount := account.NewWithUserID("2", "21", 0, "EUR")

	t.Run("happy path - transfer successful", func(t *testing.T) {
		s := setupTransferSuite()
		defer s.assertMocks(t)

		s.accRepo.On("Find", mock.Anything, originAccount.Id()).Return(originAccount, nil).Once()
		s.accRepo.On("Find", mock.Anything, destinationAccount.Id()).Return(destinationAccount, nil).Once()
		s.trxRepo.On("Transfer", mock.Anything, mock.Anything).Return(nil).Once()
		s.publisher.On("Publish", mock.Anything, mock.Anything).Return(nil).Twice()

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

	//t.Run("return 401 if user does not own the origin account", func(t *testing.T) {
	//	s := setupTransferSuite()
	//	defer s.assertMocks(t)
	//
	//	s.accRepo.On("Find", mock.Anything, mock.Anything).Return(destinationAccount, nil).Twice()
	//
	//	body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 50})
	//	require.NoError(t, err)
	//
	//	w := httptest.NewRecorder()
	//	req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//
	//	ctx := e.NewContext(req, w)
	//	ctx.Set("userId", "999")
	//
	//	err = s.ctrl.Handle(ctx)
	//	assert.NoError(t, err)
	//	assert.Equal(t, http.StatusUnauthorized, w.Code)
	//})
	//
	//t.Run("return 422 if origin account does not have enough balance", func(t *testing.T) {
	//	s := setupTransferSuite()
	//	defer s.assertMocks(t)
	//
	//	s.accRepo.On("Find", mock.Anything, originAccount.Id()).Return(originAccount, nil).Once()
	//	s.accRepo.On("Find", mock.Anything, destinationAccount.Id()).Return(destinationAccount, nil).Once()
	//
	//	body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 1500})
	//	require.NoError(t, err)
	//
	//	w := httptest.NewRecorder()
	//	req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//
	//	ctx := e.NewContext(req, w)
	//	ctx.Set("userId", "999")
	//
	//	err = s.ctrl.Handle(ctx)
	//	assert.NoError(t, err)
	//	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
	//})
	//
	//t.Run("return 404 if destination account is not found", func(t *testing.T) {
	//	s := setupTransferSuite()
	//	defer s.assertMocks(t)
	//
	//	s.accRepo.On("Find", mock.Anything, originAccount.Id()).Return(originAccount, nil).Once()
	//	s.accRepo.On("Find", mock.Anything, destinationAccount.Id()).Return(account.Account{}, errors.New("not found")).Once()
	//
	//	body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 50})
	//	require.NoError(t, err)
	//
	//	w := httptest.NewRecorder()
	//	req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//
	//	ctx := e.NewContext(req, w)
	//	ctx.Set("userId", "999")
	//
	//	err = s.ctrl.Handle(ctx)
	//	assert.NoError(t, err)
	//	assert.Equal(t, http.StatusNotFound, w.Code)
	//})
	//
	//t.Run("return 500 if transfer service fails", func(t *testing.T) {
	//	s := setupTransferSuite()
	//	defer s.assertMocks(t)
	//
	//	s.accRepo.On("Find", mock.Anything, originAccount.Id()).Return(originAccount, nil).Once()
	//	s.accRepo.On("Find", mock.Anything, destinationAccount.Id()).Return(destinationAccount, nil).Once()
	//
	//	s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(errTransfer).Once()
	//	body, err := json.Marshal(transferRequest{From: "1", To: "2", Amount: 50})
	//	require.NoError(t, err)
	//
	//	w := httptest.NewRecorder()
	//	req := httptest.NewRequest(http.MethodPost, "/transfer", bytes.NewBuffer(body))
	//	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	//
	//	ctx := e.NewContext(req, w)
	//	ctx.Set("userId", "999")
	//
	//	err = s.ctrl.Handle(ctx)
	//	assert.NoError(t, err)
	//	assert.Equal(t, http.StatusInternalServerError, w.Code)
	//})
}
