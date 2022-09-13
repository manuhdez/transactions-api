package service

import (
	"context"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type IncreaseBalanceService struct {
	repository account.Repository
}

func NewIncreaseBalanceService(repository account.Repository) IncreaseBalanceService {
	return IncreaseBalanceService{repository}
}

func (s IncreaseBalanceService) Increase(accountId string, amount float32) error {
	a, err := s.repository.Find(context.Background(), accountId)
	if err != nil {
		log.Printf("Failed to find repository: %e", err)
		return err
	}

	newBalance := a.Balance() + amount
	err = s.repository.UpdateBalance(context.Background(), accountId, newBalance)
	if err != nil {
		log.Printf("Failed to update balance: %e", err)
		return err
	}

	return nil
}
