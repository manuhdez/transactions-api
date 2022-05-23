package infra

import (
	"context"
	"database/sql"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

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

func (r AccountMysqlRepository) FindAll(ctx context.Context) ([]account.Account, error) {
	rows, err := r.db.QueryContext(ctx, "select * from accounts")
	if err != nil {
		return []account.Account{}, err
	}

	defer rows.Close()

	var accounts []AccountMysql
	for rows.Next() {
		var account AccountMysql
		if err := rows.Scan(&account.Id, &account.Balance); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	if err := rows.Err(); err != nil {
		return []account.Account{}, nil
	}

	return parseToDomainModels(accounts), nil
}

func (r AccountMysqlRepository) Find(ctx context.Context, id string) (account.Account, error) {
	row := r.db.QueryRowContext(ctx, "select * from accounts where id=?", id)

	var a AccountMysql
	if err := row.Scan(&a.Id, &a.Balance); err != nil {
		return account.Account{}, err
	}

	return a.parseToDomainModel(), nil
}
