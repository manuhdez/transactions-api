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

func (r TransactionMysqlRepository) Deposit(ctx context.Context, deposit transaction.Transaction) error {
	return r.saveTransaction(ctx, deposit)
}

func (r TransactionMysqlRepository) Withdraw(ctx context.Context, withdraw transaction.Transaction) error {
	return r.saveTransaction(ctx, withdraw)
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

func (r TransactionMysqlRepository) FindByAccount(ctx context.Context, id string) ([]transaction.Transaction, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM transactions WHERE account_id LIKE ?", id)
	if err != nil {
		log.Printf("Query failed: %e", err)
		return []transaction.Transaction{}, err
	}

	defer rows.Close()

	var tt []transaction.Transaction
	for rows.Next() {
		var t TransactionMysql
		if err := rows.Scan(&t.Id, &t.AccountId, &t.Amount, &t.Balance, &t.Type, &t.Date); err != nil {
			log.Printf("Failed to scan transaction row: %e", err)
			return []transaction.Transaction{}, err
		}
		tt = append(tt, t.ToDomainModel())
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error scanning rows: %e", err)
		return []transaction.Transaction{}, err
	}

	return tt, nil
}

func (r TransactionMysqlRepository) saveTransaction(ctx context.Context, trans transaction.Transaction) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO transactions (account_id, amount, type, balance, date) VALUES (?, ?, ?, ?, ?)",
		trans.AccountId,
		trans.Amount,
		trans.Type,
		trans.Amount,
		time.Now(),
	)

	if err != nil {
		log.Printf("Error saving %s transaction: %e", trans.Type, err)
		return err
	}

	return nil
}
