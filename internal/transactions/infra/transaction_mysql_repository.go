package infra

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type TransactionMysqlRepository struct {
	db *sql.DB
}

func NewTransactionMysqlRepository(db *sql.DB) TransactionMysqlRepository {
	return TransactionMysqlRepository{db: db}
}

func (r TransactionMysqlRepository) Deposit(ctx context.Context, transaction transaction.Transaction) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO transactions (account_id, amount, type, balance, date) VALUES (?, ?, ?, ?, ?)",
		transaction.AccountId,
		transaction.Amount,
		transaction.Type,
		transaction.Amount,
		time.Now(),
	)

	if err != nil {
		log.Printf("Error saving deposit: %e", err)
		return err
	}

	return nil
}

func (r TransactionMysqlRepository) FindAll(ctx context.Context) ([]transaction.Transaction, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM transactions")
	if err != nil {
		return []transaction.Transaction{}, err
	}

	defer rows.Close()

	var transactions []transaction.Transaction
	for rows.Next() {
		var t TransactionMysql
		if er := rows.Scan(&t.Id, &t.AccountId, &t.Amount, &t.Balance, &t.Type, &t.Date); er != nil {
			return []transaction.Transaction{}, er
		}
		transactions = append(transactions, t.ToDomainModel())
	}

	if err = rows.Err(); err != nil {
		return []transaction.Transaction{}, err
	}

	return transactions, nil
}
