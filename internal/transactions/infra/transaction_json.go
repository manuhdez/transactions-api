package infra

import (
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"log"
)

type JsonTransaction struct {
	Account  string  `json:"account_id"`
	Type     string  `json:"type"`
	Amount   float32 `json:"amount"`
	Currency string  `json:"currency"`
}

func NewJsonTransaction(t transaction.Transaction) JsonTransaction {
	log.Printf("new transaction from %v", t)
	return JsonTransaction{t.AccountId, string(t.Type), t.Amount, t.Currency}
}
