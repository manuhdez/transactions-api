package service

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type UsersRetriever struct {
	repository user.Repository
}

func NewUsersRetrieverService(repository user.Repository) UsersRetriever {
	return UsersRetriever{repository: repository}
}

func (srv UsersRetriever) Retrieve() ([]user.User, error) {
	ctx := context.Background()
	return srv.repository.All(ctx)
}
