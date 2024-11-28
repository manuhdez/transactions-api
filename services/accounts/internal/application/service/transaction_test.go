package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
)

var ErrSaveTrx = errors.New("error saving transaction")

type trxSuite struct {
	trxRepo *mocks.TransactionRepository
	srv     *TransactionService
}

func (s trxSuite) assertMocks(t *testing.T) {
	s.trxRepo.AssertExpectations(t)
}

func setupSuite() trxSuite {
	trxRepo := new(mocks.TransactionRepository)

	return trxSuite{
		trxRepo: trxRepo,
		srv:     NewTransactionService(trxRepo),
	}
}

func TestTransactionService_Deposit(t *testing.T) {
	t.Run("happy path - deposit successful", func(t *testing.T) {
		s := setupSuite()
		s.trxRepo.On("Deposit", mock.Anything, mock.Anything).Return(nil).Once()

		err := s.srv.Deposit(context.TODO(), transaction.Transaction{
			Type:      transaction.Deposit,
			AccountId: "123",
			Amount:    25,
			UserId:    "000",
		})

		s.assertMocks(t)
		assert.NoError(t, err)
	})

	t.Run("error invalid transaction type", func(t *testing.T) {
		s := setupSuite()
		err := s.srv.Deposit(context.Background(), transaction.Transaction{
			Type:      transaction.Withdrawal,
			AccountId: "123",
			Amount:    25,
		})

		s.assertMocks(t)
		assert.ErrorIs(t, err, ErrInvalidTransactionType)
	})

	t.Run("error saving deposit in repository", func(t *testing.T) {
		s := setupSuite()
		s.trxRepo.On("Deposit", mock.Anything, mock.Anything).Return(ErrSaveTrx).Once()

		err := s.srv.Deposit(context.TODO(), transaction.Transaction{
			Type:      transaction.Deposit,
			AccountId: "123",
			Amount:    25,
			UserId:    "000",
		})

		s.assertMocks(t)
		assert.ErrorIs(t, err, ErrSaveTrx)
	})
}

func TestTransactionService_Withdraw(t *testing.T) {
	t.Run("happy path - withdraw successful", func(t *testing.T) {
		s := setupSuite()
		s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(nil).Once()

		err := s.srv.Withdraw(context.TODO(), transaction.Transaction{
			Type:      transaction.Withdrawal,
			AccountId: "123",
			Amount:    25,
			UserId:    "000",
		})

		s.assertMocks(t)
		assert.NoError(t, err)
	})

	t.Run("error invalid transaction type", func(t *testing.T) {
		s := setupSuite()

		err := s.srv.Withdraw(context.Background(), transaction.Transaction{
			Type:      transaction.Deposit,
			AccountId: "123",
			Amount:    25,
		})

		s.assertMocks(t)
		assert.ErrorIs(t, err, ErrInvalidTransactionType)
	})

	t.Run("error saving withdraw in repository", func(t *testing.T) {
		s := setupSuite()
		s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(ErrSaveTrx).Once()

		err := s.srv.Withdraw(context.TODO(), transaction.Transaction{
			Type:      transaction.Withdrawal,
			AccountId: "123",
			Amount:    25,
			UserId:    "000",
		})

		s.assertMocks(t)
		assert.ErrorIs(t, err, ErrSaveTrx)
	})
}

func TestTransactionService_Transfer(t *testing.T) {
	var (
		userID        = "111"
		fromAccountID = "001"
		toAccountID   = "002"
	)

	t.Run("happy path - transfer successful", func(t *testing.T) {
		s := setupSuite()
		s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(nil).Once()
		s.trxRepo.On("Deposit", mock.Anything, mock.Anything).Return(nil).Once()

		trf := transaction.NewTransfer(userID, fromAccountID, toAccountID, 120)
		err := s.srv.Transfer(context.TODO(), trf)
		assert.NoError(t, err)
		s.assertMocks(t)
	})

	t.Run("error withdrawing from origin account", func(t *testing.T) {
		s := setupSuite()
		s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(ErrSaveTrx).Once()

		trf := transaction.NewTransfer(userID, fromAccountID, toAccountID, 120)
		err := s.srv.Transfer(context.TODO(), trf)
		assert.ErrorIs(t, err, ErrSaveTrx)
		s.assertMocks(t)
	})

	t.Run("error depositing into destination account", func(t *testing.T) {
		s := setupSuite()
		s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(nil).Once()
		s.trxRepo.On("Deposit", mock.Anything, mock.Anything).Return(ErrSaveTrx).Once()

		trf := transaction.NewTransfer(userID, fromAccountID, toAccountID, 120)
		err := s.srv.Transfer(context.TODO(), trf)
		assert.ErrorIs(t, err, ErrSaveTrx)
		s.assertMocks(t)
	})

	t.Run("error transfer amount too low", func(t *testing.T) {
		s := setupSuite()
		err := s.srv.Transfer(
			context.TODO(),
			transaction.NewTransfer(userID, fromAccountID, toAccountID, 8),
		)
		assert.ErrorIs(t, err, ErrTransferAmountTooLow)
	})

	t.Run("error transfer amount too high", func(t *testing.T) {
		s := setupSuite()
		err := s.srv.Transfer(
			context.TODO(),
			transaction.NewTransfer(userID, fromAccountID, toAccountID, 120_000),
		)
		assert.ErrorIs(t, err, ErrTransferAmountTooHigh)
	})
}
