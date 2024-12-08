package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	"github.com/manuhdez/transactions-api/shared/domain"
)

type transferSuite struct {
	trxRepo *mocks.TransactionRepository
	bus     *mocks.EventBus
}

func setup() transferSuite {
	trxRepo := new(mocks.TransactionRepository)
	bus := new(mocks.EventBus)

	return transferSuite{
		trxRepo: trxRepo,
		bus:     bus,
	}
}

func (s transferSuite) assertMocks(t *testing.T) {
	s.trxRepo.AssertExpectations(t)
	s.bus.AssertExpectations(t)
}

func TestTransfer_Transfer(t *testing.T) {
	userID := domain.NewID("111")
	origin := account.NewWithUserID("001", userID, 100, "EUR")
	destination := account.NewWithUserID("002", domain.NewID("025"), 100, "EUR")

	t.Run("happy path - should transfer money between account", func(t *testing.T) {
		s := setup()
		s.trxRepo.On("Transfer", mock.Anything, mock.Anything).Return(nil).Once()
		s.bus.On("Publish", mock.Anything, mock.Anything).Return(nil).Twice()

		srv := NewTransferService(s.trxRepo, s.bus)
		trx := transaction.CreateTransfer(userID, origin, destination, 20)
		err := srv.Transfer(context.TODO(), trx)
		assert.NoError(t, err)
		s.assertMocks(t)
	})

	t.Run("should return error if user is not owner of the origin account", func(t *testing.T) {
		s := setup()

		srv := NewTransferService(s.trxRepo, s.bus)
		err := srv.Transfer(context.TODO(), transaction.NewTransfer(userID, destination, origin, 20))
		assert.ErrorIs(t, err, ErrUnauthorizedTransaction)
		s.assertMocks(t)
	})

	t.Run("should error if amount is bigger than the maximum allowed transfer amount", func(t *testing.T) {
		s := setup()
		srv := NewTransferService(s.trxRepo, s.bus)
		err := srv.Transfer(context.TODO(), transaction.NewTransfer(userID, origin, destination, transaction.MaxTransferAmount+1))
		assert.ErrorIs(t, err, transaction.ErrTransferAmountTooBig)
		s.assertMocks(t)
	})

	t.Run("should error if amount is smaller than the minimum allowed transfer amount", func(t *testing.T) {
		s := setup()
		srv := NewTransferService(s.trxRepo, s.bus)
		err := srv.Transfer(context.TODO(), transaction.NewTransfer(userID, origin, destination, 0))
		assert.ErrorIs(t, err, transaction.ErrTransferAmountTooSmall)
		s.assertMocks(t)
	})

	t.Run("should error if origin account has insufficient balance", func(t *testing.T) {
		s := setup()
		emptyOrigin := account.NewWithUserID("1", userID, 0, "EUR")

		srv := NewTransferService(s.trxRepo, s.bus)
		err := srv.Transfer(context.TODO(), transaction.NewTransfer(userID, emptyOrigin, destination, 100))
		assert.ErrorIs(t, err, transaction.ErrInsufficientBalance)
		s.assertMocks(t)
	})

}
