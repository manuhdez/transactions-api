package service_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/transactions/app/service"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
)

type suite struct {
	trxRepo  *mocks.TransactionMockRepository
	eventBus *mocks.EventBus
	withdraw service.Withdraw
}

func newTestSuite() suite {
	trxRepo := new(mocks.TransactionMockRepository)
	eventBus := new(mocks.EventBus)

	return suite{
		trxRepo:  trxRepo,
		eventBus: eventBus,
		withdraw: service.NewWithdrawService(trxRepo, eventBus),
	}
}

func (s *suite) assertMocks(t *testing.T) {
	s.trxRepo.AssertExpectations(t)
	s.eventBus.AssertExpectations(t)
}

func TestWithDraw_Invoke_Success(t *testing.T) {
	s := newTestSuite()
	defer s.assertMocks(t)

	s.trxRepo.On("Withdraw", context.Background(), mock.Anything).Return(nil).Once()
	s.eventBus.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()

	err := s.withdraw.Invoke(
		context.Background(),
		transaction.NewTransaction(transaction.Withdrawal, "1", 125.5, "EUR"),
	)
	assert.NoError(t, err)
}

func TestWithdraw_Invalid_Trx(t *testing.T) {
	s := newTestSuite()
	defer s.assertMocks(t)

	ctx := context.Background()
	trx := transaction.NewTransaction("invalid transaction type", "1", 125.5, "EUR")
	err := s.withdraw.Invoke(ctx, trx)
	assert.ErrorIs(t, err, service.ErrInvalidTransactionType)
}

func TestWithDraw_Invoke_Repo_Fail(t *testing.T) {
	s := newTestSuite()
	defer s.assertMocks(t)

	expected := errors.New("could not create the withdraw")
	s.trxRepo.On("Withdraw", context.Background(), mock.Anything).Return(expected)

	err := s.withdraw.Invoke(context.Background(), transaction.NewTransaction(transaction.Withdrawal, "23", 33253, "EUR"))
	assert.ErrorIs(t, err, expected)
}
