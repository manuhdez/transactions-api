package infra

import "github.com/manuhdez/transactions-api/internal/accounts/domain/account"

type AccountMysql struct {
	Id       string  `mysql:"id"`
	Balance  float32 `mysql:"balance"`
	Currency string  `mysql:"currency"`
}

func (a AccountMysql) parseToDomainModel() account.Account {
	return account.New(a.Id, a.Balance, a.Currency)
}

func parseToDomainModels(list []AccountMysql) []account.Account {
	var accounts []account.Account
	for _, item := range list {
		accounts = append(accounts, item.parseToDomainModel())
	}
	return accounts
}
