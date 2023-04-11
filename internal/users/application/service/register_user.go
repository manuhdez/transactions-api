package service

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/users/domain/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type RegisterUser struct {
	repository user.Repository
	hasher     service.HashService
}

func NewRegisterUserService(repository user.Repository, hasher service.HashService) RegisterUser {
	return RegisterUser{repository, hasher}
}

func (srv RegisterUser) Register(user user.User) error {
	ctx := context.Background()

	hashedPassword, err := srv.hasher.Hash(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return srv.repository.Save(ctx, user)
}
