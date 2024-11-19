package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
)

var (
	ErrSaveTrx         = errors.New("error saving transaction")
	ErrAccountNotFound = errors.New("account not found")
)

type trxSuite struct {
	trxRepo *mocks.TransactionMockRepository
	accRepo *mocks.AccountMockRepository
	srv     *TransactionService
}

func (s trxSuite) assertMocks(t *testing.T) {
	s.trxRepo.AssertExpectations(t)
	s.accRepo.AssertExpectations(t)
}

func setupSuite() trxSuite {
	trxRepo := new(mocks.TransactionMockRepository)
	accRepo := new(mocks.AccountMockRepository)

	return trxSuite{
		trxRepo: trxRepo,
		accRepo: accRepo,
		srv:     NewTransactionService(trxRepo, accRepo),
	}
}

func TestTransactionService_Deposit(t *testing.T) {
	t.Run("happy path - deposit successful", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{UserId: "000"}, nil)
		s.trxRepo.On("Deposit", mock.Anything, mock.Anything).Return(nil).Once()
		//s.eventBus.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()

		//ctx := context.WithValue(context.Background(), "userId", "000")
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
		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{UserId: "000"}, nil)
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

	t.Run("error account not found", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{}, ErrAccountNotFound)

		err := s.srv.Deposit(context.Background(), transaction.Transaction{
			Type:      transaction.Deposit,
			AccountId: "123",
			Amount:    25,
		})

		s.assertMocks(t)
		assert.ErrorIs(t, err, ErrAccountNotFound)
	})

	t.Run("error unauthorized transaction", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{UserId: "000"}, nil)

		err := s.srv.Deposit(context.TODO(), transaction.Transaction{
			Type:      transaction.Deposit,
			AccountId: "123",
			Amount:    25,
			UserId:    "111",
		})

		s.assertMocks(t)
		assert.ErrorIs(t, err, ErrUnauthorized)
	})
}

func TestTransactionService_Withdraw(t *testing.T) {
	t.Run("happy path - withdraw successful", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{UserId: "000"}, nil)
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
		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{UserId: "000"}, nil)
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

	t.Run("error account not found", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{}, ErrAccountNotFound)

		err := s.srv.Withdraw(context.Background(), transaction.Transaction{
			Type:      transaction.Withdrawal,
			AccountId: "123",
			Amount:    25,
		})

		s.assertMocks(t)
		assert.ErrorIs(t, err, ErrAccountNotFound)
	})

	t.Run("error unauthorized transaction", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, mock.Anything).Return(account.Account{UserId: "111"}, nil)

		err := s.srv.Withdraw(context.TODO(), transaction.Transaction{
			Type:      transaction.Withdrawal,
			AccountId: "123",
			Amount:    25,
			UserId:    "000",
		})

		s.assertMocks(t)
		assert.ErrorIs(t, err, ErrUnauthorized)
	})
}

func TestTransactionService_Transfer(t *testing.T) {
	var (
		userId        = "111"
		fromAccountId = "001"
		toAccountId   = "002"
	)

	t.Run("happy path - transfer successful", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, fromAccountId).Return(account.Account{UserId: userId}, nil).Once()
		s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(nil).Once()
		s.trxRepo.On("Deposit", mock.Anything, mock.Anything).Return(nil).Once()

		trf := transaction.NewTransfer(userId, fromAccountId, toAccountId, 120)
		err := s.srv.Transfer(context.TODO(), trf)
		assert.NoError(t, err)
		s.assertMocks(t)
	})

	t.Run("error unauthorized access to origin account", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, fromAccountId).Return(account.Account{UserId: "333"}, nil).Once()

		trf := transaction.NewTransfer(userId, fromAccountId, toAccountId, 120)
		err := s.srv.Transfer(context.TODO(), trf)
		assert.ErrorIs(t, err, ErrUnauthorized)
		s.assertMocks(t)
	})

	t.Run("error withdrawing from origin account", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, fromAccountId).Return(account.Account{UserId: userId}, nil).Once()
		s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(ErrSaveTrx).Once()

		trf := transaction.NewTransfer(userId, fromAccountId, toAccountId, 120)
		err := s.srv.Transfer(context.TODO(), trf)
		assert.ErrorIs(t, err, ErrSaveTrx)
		s.assertMocks(t)
	})

	t.Run("error depositing into destination account", func(t *testing.T) {
		s := setupSuite()
		s.accRepo.On("FindById", mock.Anything, fromAccountId).Return(account.Account{UserId: userId}, nil).Once()
		s.trxRepo.On("Withdraw", mock.Anything, mock.Anything).Return(nil).Once()
		s.trxRepo.On("Deposit", mock.Anything, mock.Anything).Return(ErrSaveTrx).Once()

		trf := transaction.NewTransfer(userId, fromAccountId, toAccountId, 120)
		err := s.srv.Transfer(context.TODO(), trf)
		assert.ErrorIs(t, err, ErrSaveTrx)
		s.assertMocks(t)
	})
}
