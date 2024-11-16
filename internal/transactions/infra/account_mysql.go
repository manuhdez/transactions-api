package infra

import "github.com/manuhdez/transactions-api/internal/transactions/domain/account"

type AccountMysql struct {
	Id     string `mysql:"id"`
	UserId string `mysql:"user_id"`
}

// ToDomainModel converts a AccountMysql to a domain account.Account
func (a AccountMysql) ToDomainModel() account.Account {
	return account.NewAccount(a.Id, a.UserId)
}
