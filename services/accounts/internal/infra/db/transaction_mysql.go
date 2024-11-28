package db

import (
	"time"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
)

type TransactionMysql struct {
	Id        int              `mysql:"id"`
	AccountId string           `mysql:"account_id"`
	UserId    string           `mysql:"user_id"`
	Amount    float32          `mysql:"amount"`
	Balance   float32          `mysql:"balance"`
	Type      transaction.Type `mysql:"type"`
	Date      time.Time        `mysql:"date"`
}

func (t TransactionMysql) TableName() string {
	return "transactions"
}

func (t TransactionMysql) ToDomainModel() transaction.Transaction {
	trx := transaction.NewTransaction(t.Type, t.AccountId, t.UserId, t.Amount)
	trx.Date = t.Date
	return trx
}
