package infra

import (
	"time"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type JsonTransaction struct {
	Account string    `json:"account_id"`
	Type    string    `json:"type"`
	Amount  float32   `json:"amount"`
	Date    time.Time `json:"date"`
}

func NewJsonTransaction(t transaction.Transaction) JsonTransaction {
	return JsonTransaction{
		Account: t.AccountId,
		Type:    string(t.Type),
		Amount:  t.Amount,
		Date:    t.Date,
	}
}
