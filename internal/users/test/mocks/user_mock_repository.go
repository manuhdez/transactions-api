package mocks

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type UserMockRepository struct {
	user.Repository
	Err error
}

func (repo UserMockRepository) All(_ context.Context) ([]user.User, error) {
	return []user.User{}, repo.Err
}

func (repo UserMockRepository) Save(_ context.Context, _ user.User) error {
	return repo.Err
}

func (repo UserMockRepository) FindByEmail(_ context.Context, email string) (user.User, error) {
	return user.User{Id: "1", Email: email, Password: "test-password"}, repo.Err
}
