package infra

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type TransactionMysqlRepository struct {
	db *sql.DB
}

func NewTransactionMysqlRepository(db *sql.DB) TransactionMysqlRepository {
	return TransactionMysqlRepository{db: db}
}

func (r TransactionMysqlRepository) Deposit(ctx context.Context, transaction transaction.Transaction) error {
	fmt.Println("Deposit", transaction)
	return nil
}
