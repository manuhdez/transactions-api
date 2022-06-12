package infra

import (
	"context"
	"database/sql"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type TransactionMysqlRepository struct {
	db *sql.DB
}

func NewTransactionMysqlRepository(db *sql.DB) TransactionMysqlRepository {
	return TransactionMysqlRepository{db: db}
}

func (r TransactionMysqlRepository) Deposit(ctx context.Context, t transaction.Transaction) error {
	query := "insert into transactions (account_id, amount, currency, type) values (?, ?, ?, ?)"
	_, err := r.db.ExecContext(ctx, query, t.AccountId, t.Amount, t.Currency, t.Type)
	if err != nil {
		return err
	}

	return nil
}
