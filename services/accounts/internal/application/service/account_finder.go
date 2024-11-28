package service

import (
	"context"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
)

type AccountFinder struct {
	repo account.Repository
}

func NewAccountFinder(r account.Repository) *AccountFinder {
	return &AccountFinder{repo: r}
}

// Find returns an account found by id or an error if not found
func (f AccountFinder) Find(ctx context.Context, accountID string) (account.Account, error) {
	log.Printf("[AccountFinder:Find][accountID:%s]", accountID)

	acc, err := f.repo.Find(ctx, accountID)
	if err != nil {
		return account.Account{}, fmt.Errorf("[AccountFinder:Find][accountID:%s]%w", accountID, err)
	}

	return acc, nil
}
