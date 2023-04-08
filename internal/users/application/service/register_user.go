package service

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type RegisterUser struct {
	repository user.Repository
}

func NewRegisterUserService(repository user.Repository) RegisterUser {
	return RegisterUser{repository}
}

func (srv RegisterUser) Register(user user.User) error {
	ctx := context.Background()
	return srv.repository.Save(ctx, user)
}
