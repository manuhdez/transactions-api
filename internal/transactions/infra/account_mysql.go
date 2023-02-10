package infra

import "github.com/manuhdez/transactions-api/internal/transactions/domain/account"

type AccountMysql struct {
	Id string `mysql:"id"`
}

func (a AccountMysql) ToDomainModel() account.Account {
	return account.NewAccount(a.Id)
}
