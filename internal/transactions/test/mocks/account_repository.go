package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
)

type AccountMockRepository struct {
	mock.Mock
}

func (m *AccountMockRepository) Save(ctx context.Context, acc account.Account) error {
	args := m.Called(ctx, acc)
	return args.Error(0)
}

func (m *AccountMockRepository) FindById(ctx context.Context, id string) (account.Account, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(account.Account), args.Error(1)
}
