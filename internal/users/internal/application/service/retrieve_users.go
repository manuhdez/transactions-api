package service

import (
	"context"

	user2 "github.com/manuhdez/transactions-api/internal/users/internal/domain/user"
)

type UsersRetriever struct {
	repository user2.Repository
}

func NewUsersRetrieverService(repository user2.Repository) UsersRetriever {
	return UsersRetriever{repository: repository}
}

func (srv UsersRetriever) Retrieve() ([]user2.User, error) {
	ctx := context.Background()
	return srv.repository.All(ctx)
}
