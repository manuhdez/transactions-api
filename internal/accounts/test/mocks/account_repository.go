package mocks

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/stretchr/testify/mock"
)

type AccountMockRepository struct {
	mock.Mock
}

func (m *AccountMockRepository) Create(a account.Account) error {
	args := m.Called(a)
	return args.Error(0)
}

func (m *AccountMockRepository) FindAll(c context.Context) ([]account.Account, error) {
	args := m.Called(c)
	return args.Get(0).([]account.Account), args.Error(1)
}

func (m *AccountMockRepository) Find(c context.Context, id string) (account.Account, error) {
	args := m.Called(c, id)
	return args.Get(0).(account.Account), args.Error(1)
}
