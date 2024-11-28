package service

import (
	"context"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
)

type DecreaseBalance struct {
	repository account.Repository
}

func NewDecreaseBalanceService(repository account.Repository) DecreaseBalance {
	return DecreaseBalance{repository}
}

func (srv DecreaseBalance) Decrease(accountId string, amount float32) error {
	acc, err := srv.repository.Find(context.Background(), accountId)
	if err != nil {
		log.Printf("Account with id %s was not found", accountId)
		return err
	}

	updatedBalance := acc.Balance() - amount
	err = srv.repository.UpdateBalance(context.Background(), accountId, updatedBalance)
	if err != nil {
		log.Printf("Couldn't update balance on account %s", accountId)
		return err
	}

	return nil
}
