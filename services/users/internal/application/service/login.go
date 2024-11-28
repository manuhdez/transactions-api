package service

import (
	"context"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/users/internal/domain/service"
	user2 "github.com/manuhdez/transactions-api/internal/users/internal/domain/user"
)

type LoginService struct {
	repository user2.Repository
	hasher     domain_service.HashService
}

func NewLoginService(repo user2.Repository, hasher domain_service.HashService) LoginService {
	return LoginService{repo, hasher}
}

func (srv LoginService) Login(email, password string) (*user2.User, error) {
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
