package infra

import (
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type TransactionMysql struct {
	Id        int              `mysql:"id"`
	AccountId string           `mysql:"account_id"`
	Amount    float32          `mysql:"amount"`
	Balance   float32          `mysql:"balance"`
	Type      transaction.Type `mysql:"type"`
	Date      []uint8          `mysql:"date"`
}

func (t TransactionMysql) ToDomainModel() transaction.Transaction {
	return transaction.NewTransaction(t.Type, t.AccountId, t.Amount, "EUR")
}
