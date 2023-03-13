package infra

import (
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type JsonTransaction struct {
	Account string  `json:"account_id"`
	Type    string  `json:"type"`
	Amount  float32 `json:"amount"`
}

func NewJsonTransaction(t transaction.Transaction) JsonTransaction {
	return JsonTransaction{t.AccountId, string(t.Type), t.Amount}
}
