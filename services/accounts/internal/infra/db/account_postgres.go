package db

import (
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/shared/domain"
)

type AccountPostgres struct {
	Id       string  `mysql:"id"`
	UserId   string  `mysql:"user_id" default:"NULL"`
	Balance  float32 `mysql:"balance"`
	Currency string  `mysql:"currency"`
}

func (a AccountPostgres) TableName() string {
	return "accounts"
}

func (a AccountPostgres) parseToDomainModel() account.Account {
	return account.NewWithUserID(a.Id, domain.NewID(a.UserId), a.Balance, a.Currency)
}

func parseToDomainModels(list []AccountPostgres) []account.Account {
	var accounts []account.Account
	for _, item := range list {
		accounts = append(accounts, item.parseToDomainModel())
	}
	return accounts
}
