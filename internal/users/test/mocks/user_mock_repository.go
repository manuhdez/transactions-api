package mocks

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type UserMockRepository struct {
	user.Repository
	Err error
}

func (repo UserMockRepository) Save(_ context.Context, _ user.User) error {
	return repo.Err
}
