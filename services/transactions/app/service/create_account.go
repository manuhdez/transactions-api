package service

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
)

type CreateAccount struct {
	repository account.Repository
}

func NewCreateAccountService(repo account.Repository) CreateAccount {
	return CreateAccount{repo}
}

func (srv CreateAccount) Create(ctx context.Context, acc account.Account) error {
	err := srv.repository.Save(ctx, acc)
	if err != nil {
		return err
	}
	return nil
}
