package service

import (
	"context"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
)

type AccountsFinder struct {
	repository account.Repository
}

func NewFindAccountService(repo account.Repository) AccountsFinder {
	return AccountsFinder{repo}
}

func (f AccountsFinder) FindById(ctx context.Context, id string) (account.Account, error) {
	log.Printf("[AccountsFinder:FindById][id:%s]", id)
	return f.repository.Find(ctx, id)
}

func (f AccountsFinder) FindByUserId(ctx context.Context, userId string) ([]account.Account, error) {
	log.Printf("[AccountsFinder:FindByUserId][userId:%s]", userId)
	return f.repository.GetByUserId(ctx, userId)
}
