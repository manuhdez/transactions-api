package mocks

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/stretchr/testify/mock"
)

type TransactionMockRepository struct {
	mock.Mock
}

func (m *TransactionMockRepository) Deposit(ctx context.Context, t transaction.Transaction) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *TransactionMockRepository) Withdraw(ctx context.Context, t transaction.Transaction) error {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *TransactionMockRepository) FindAll(ctx context.Context) ([]transaction.Transaction, error) {
	args := m.Called(ctx)
	return args.Get(0).([]transaction.Transaction), args.Error(1)
}
