package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
)

func NewTestSuite() (*mocks.TransactionMockRepository, *mocks.EventBus, Withdraw) {
	repository := new(mocks.TransactionMockRepository)
	bus := new(mocks.EventBus)
	service := NewWithdrawService(repository, bus)

	return repository, bus, service
}

func TestWithdrawService(t *testing.T) {

	t.Run(
		"Should create a new withdraw", func(t *testing.T) {
			repository, _, service := NewTestSuite()

			repository.On("Withdraw", context.Background(), mock.Anything).Return(nil)
			withdraw := transaction.NewTransaction(transaction.Withdrawal, "1", 125.5, "EUR")

			err := service.Invoke(context.Background(), withdraw)
			assert.NoError(t, err)
		},
	)

	t.Run(
		"Should return an error when creating a new withdraw", func(t *testing.T) {
			repository, _, service := NewTestSuite()

			expected := errors.New("could not create the withdraw")
			repository.On("Withdraw", context.Background(), mock.Anything).Return(expected)
			withdraw := transaction.NewTransaction(transaction.Withdrawal, "23", 33253, "EUR")

			res := service.Invoke(context.Background(), withdraw)

			if assert.Error(t, res) {
				assert.Equal(t, expected, res)
			}
		},
	)
}
