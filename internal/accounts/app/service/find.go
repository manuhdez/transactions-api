package service

import (
	"context"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type FindAccountService struct {
	repository account.Repository
}

func NewFindAccountService(repo account.Repository) FindAccountService {
	return FindAccountService{repo}
}

func (s FindAccountService) Find(ctx context.Context, id string) (account.Account, error) {
	return s.repository.Find(ctx, id)
}
