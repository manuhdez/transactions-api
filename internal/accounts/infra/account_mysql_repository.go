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
	_, err := r.db.Exec("INSERT INTO accounts (id, balance, currency) VALUES ($1, $2, $3)", a.Id(), a.Balance(), a.Currency())
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
		var acc AccountMysql
		if err = rows.Scan(&acc.Id, &acc.Balance, &acc.Currency); err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}
	if err = rows.Err(); err != nil {
		return []account.Account{}, nil
	}

	return parseToDomainModels(accounts), nil
}

func (r AccountMysqlRepository) Find(ctx context.Context, id string) (account.Account, error) {
	row := r.db.QueryRowContext(ctx, "select * from accounts where id=$1", id)

	var a AccountMysql
	err := row.Scan(&a.Id, &a.Balance, &a.Currency)

	// this check was necessary to avoid nil pointer error
	// TODO: solve nil pointer error and remove this check
	if (a != AccountMysql{}) {
		return a.parseToDomainModel(), nil
	}

	// If the account the user is looking for does not exist, it should not throw an error
	if err.Error() == "sql: no rows in result set" {
		return account.Account{}, nil
	}

	if err != nil {
		return account.Account{}, err
	}

	return a.parseToDomainModel(), nil
}

func (r AccountMysqlRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "delete from accounts where id=$1", id)
	return err
}

func (r AccountMysqlRepository) UpdateBalance(ctx context.Context, id string, newBalance float32) error {
	_, err := r.db.ExecContext(ctx, "update accounts set balance = $1 where id = $2", newBalance, id)
	return err
}
