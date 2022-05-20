package mocks

import (
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
