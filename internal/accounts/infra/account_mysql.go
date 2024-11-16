package infra

import (
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/shared/domain"
)

type AccountMysql struct {
	Id       string  `mysql:"id"`
	UserId   string  `mysql:"user_id" default:"NULL"`
	Balance  float32 `mysql:"balance"`
	Currency string  `mysql:"currency"`
}

func (a AccountMysql) TableName() string {
	return "accounts"
}

func (a AccountMysql) parseToDomainModel() account.Account {
	return account.NewWithUserID(a.Id, domain.NewID(a.UserId), a.Balance, a.Currency)
}

func parseToDomainModels(list []AccountMysql) []account.Account {
	var accounts []account.Account
	for _, item := range list {
		accounts = append(accounts, item.parseToDomainModel())
	}
	return accounts
}
