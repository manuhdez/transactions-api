package infra

import (
	"database/sql"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type AccountMysql struct {
	Id      string  `mysql:"id"`
	Balance float32 `mysql:"balance"`
}

func (a AccountMysql) ToDomainAccount() account.Account {
	return account.New(a.Id, a.Balance)
}

type AccountMysqlRepository struct {
	db *sql.DB
}

func NewAccountMysqlRepository(db *sql.DB) AccountMysqlRepository {
	return AccountMysqlRepository{db}
}

func (r AccountMysqlRepository) Create(a account.Account) error {
	_, err := r.db.Exec("INSERT INTO accounts (id, balance) VALUES (?, ?)", a.Id(), a.Balance())
	return err
}
