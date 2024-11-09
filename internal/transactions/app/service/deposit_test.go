package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
)

type depositSuite struct {
	trxRepo  *mocks.TransactionMockRepository
	eventBus *mocks.EventBus
	deposit  service.Deposit
}

func newDepositSuite() *depositSuite {
	repo := new(mocks.TransactionMockRepository)
	bus := new(mocks.EventBus)

	return &depositSuite{
		trxRepo:  repo,
		eventBus: bus,
		deposit:  service.NewDepositService(repo, bus),
	}
}

func (s *depositSuite) assertMocks(t *testing.T) {
	s.trxRepo.AssertExpectations(t)
	s.eventBus.AssertExpectations(t)
}

func TestDeposit_Success(t *testing.T) {
	s := newDepositSuite()

	ctx := context.Background()
	trx := transaction.NewTransaction(transaction.Deposit, "user-id", 100, "EUR")

	s.trxRepo.On("Deposit", ctx, trx).Return(nil).Once()
	s.eventBus.On("Publish", ctx, mock.Anything).Return(nil).Once()
	err := s.deposit.Invoke(ctx, trx)
	assert.NoError(t, err)
}

func TestDeposit_Deposit_Error(t *testing.T) {
	s := newDepositSuite()
	defer s.assertMocks(t)

	ctx := context.Background()
	trx := transaction.NewTransaction(transaction.Deposit, "user-id", 100, "EUR")

	err := fmt.Errorf("error saving deposit transaction")
	s.trxRepo.On("Deposit", ctx, trx).Return(err).Once()

	assert.Error(t, s.deposit.Invoke(ctx, trx))
}

func TestDeposit_Publish_Error(t *testing.T) {
	s := newDepositSuite()
	defer s.assertMocks(t)

	ctx := context.Background()
	trx := transaction.NewTransaction(transaction.Deposit, "user-id", 100, "EUR")

	err := fmt.Errorf("[err: failed publishing deposit transaction event]")
	s.trxRepo.On("Deposit", ctx, trx).Return(nil).Once()
	s.eventBus.On("Publish", ctx, mock.Anything).Return(err).Once()

	assert.NoError(t, s.deposit.Invoke(ctx, trx), err)
}

func TestDeposit_Invalid_Transaction(t *testing.T) {
	s := newDepositSuite()
	defer s.assertMocks(t)

	ctx := context.Background()
	trx := transaction.NewTransaction("invalid-type", "user-id", 100, "EUR")
	assert.ErrorIs(t, s.deposit.Invoke(ctx, trx), service.ErrInvalidTransactionType)
}
