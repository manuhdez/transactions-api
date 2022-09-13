package service

import (
	"fmt"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type IncreaseBalanceService struct {
	repository account.Repository
}

func NewIncreaseBalanceService(repository account.Repository) IncreaseBalanceService {
	return IncreaseBalanceService{repository}
}

func (s IncreaseBalanceService) Increase(accountId string, amount float32) error {
	fmt.Printf("Increase account %s balance by %f", accountId, amount)
	return nil
}
