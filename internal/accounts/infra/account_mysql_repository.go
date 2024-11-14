package infra

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type AccountMysqlRepository struct {
	db *sql.DB
}

func NewAccountMysqlRepository(db *sql.DB) AccountMysqlRepository {
	return AccountMysqlRepository{db}
}

func (r AccountMysqlRepository) Create(a account.Account) error {
	_, err := r.db.Exec("INSERT INTO accounts (id, user_id, balance, currency) VALUES ($1, $2, $3, $4)", a.Id(), a.UserId.String(), a.Balance(), a.Currency())
	return err
}

func (r AccountMysqlRepository) Find(ctx context.Context, id string) (account.Account, error) {
	row := r.db.QueryRowContext(ctx, "select id, user_id, balance, currency from accounts where id=$1", id)

	var a AccountMysql
	err := row.Scan(&a.Id, &a.UserId, &a.Balance, &a.Currency)

	// this check was necessary to avoid nil pointer error
	// TODO: solve nil pointer error and remove this check
	if (a != AccountMysql{}) {
		return a.parseToDomainModel(), nil
	}

	// If the account the user is looking for does not exist, it should not throw an error
	if errors.Is(err, sql.ErrNoRows) {
		return account.Account{}, nil
	}

	if err != nil {
		return account.Account{}, err
	}

	return a.parseToDomainModel(), nil
}

// GetByUserId returns the list of accounts for a given user
func (r AccountMysqlRepository) GetByUserId(ctx context.Context, userId string) ([]account.Account, error) {
	var accounts []AccountMysql
	rows, err := r.db.QueryContext(ctx, "select id, user_id, balance, currency from accounts where user_id=$1", userId)
	if err != nil {
		log.Printf("[AccountMysqlRepository:GetByUserId][err: %s]", err)
		return nil, fmt.Errorf("[AccountMysqlRepository:GetByUserId][err: database error]")
	}

	defer rows.Close()
	for rows.Next() {
		var acc AccountMysql
		if err = rows.Scan(&acc.Id, &acc.UserId, &acc.Balance, &acc.Currency); err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}
	if err = rows.Err(); err != nil {
		return []account.Account{}, nil
	}

	return parseToDomainModels(accounts), nil
}

func (r AccountMysqlRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.ExecContext(ctx, "delete from accounts where id=$1", id)
	return err
}

func (r AccountMysqlRepository) UpdateBalance(ctx context.Context, id string, newBalance float32) error {
	_, err := r.db.ExecContext(ctx, "update accounts set balance = $1 where id = $2", newBalance, id)
	return err
}
