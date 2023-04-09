package service

import (
	"context"
	"fmt"
	"log"

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
		log.Printf("error fetching user: %e", err)
		return nil, fmt.Errorf("Invalid authentication data")
	}

	if usr.Password != password {
		return nil, fmt.Errorf("Invalid authentication data")
	}

	return &usr, nil
}
