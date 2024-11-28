package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
)

type AccountMockRepository struct {
	mock.Mock
}

func (m *AccountMockRepository) Create(ctx context.Context, a account.Account) error {
	args := m.Called(ctx, a)
	return args.Error(0)
}

func (m *AccountMockRepository) Find(c context.Context, id string) (account.Account, error) {
	args := m.Called(c, id)
	return args.Get(0).(account.Account), args.Error(1)
}

func (m *AccountMockRepository) Delete(c context.Context, id string) error {
	args := m.Called(c, id)
	return args.Error(0)
}

func (m *AccountMockRepository) UpdateBalance(ctx context.Context, id string, balance float32) error {
	args := m.Called(ctx, id, balance)
	return args.Error(0)
}

func (m *AccountMockRepository) GetByUserId(c context.Context, id string) ([]account.Account, error) {
	args := m.Called(c, id)
	return args.Get(0).([]account.Account), args.Error(1)
}
