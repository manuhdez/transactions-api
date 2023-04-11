package service

import (
	"context"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/users/domain/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type LoginService struct {
	repository user.Repository
	hasher     service.HashService
}

func NewLoginService(repo user.Repository, hasher service.HashService) LoginService {
	return LoginService{repo, hasher}
}

func (srv LoginService) Login(email, password string) (*user.User, error) {
	ctx := context.Background()
	usr, err := srv.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("invalid authentication data")
	}

	if err = srv.hasher.Compare(usr.Password, password); err != nil {
		return nil, fmt.Errorf("invalid authentication data")
	}

	return &usr, nil
}
