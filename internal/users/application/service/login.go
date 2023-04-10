package service

import (
	"context"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type LoginService struct {
	repository user.Repository
}

func NewLoginService(repo user.Repository) LoginService {
	return LoginService{repo}
}

func (srv LoginService) Login(email, password string) (*user.User, error) {
	ctx := context.Background()
	usr, err := srv.repository.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("invalid authentication data")
	}
	if usr.Password != password {
		return nil, fmt.Errorf("invalid authentication data")
	}

	return &usr, nil
}
