package event

import (
	"encoding/json"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

type DepositCreated struct {
	body []byte
}

type DepositCreatedBody struct {
	Type    string  `json:"type"`
	Account string  `json:"account"`
	Amount  float32 `json:"amount"`
}

func NewDepositCreated(trx transaction.Transaction) DepositCreated {
	body := DepositCreatedBody{string(DepositCreatedType), trx.AccountId, trx.Amount}
	b, _ := json.Marshal(body)
	return DepositCreated{b}
}

var DepositCreatedType Type = "event.transactions.deposit_created"

func (d DepositCreated) Type() Type {
	return DepositCreatedType
}

func (d DepositCreated) Body() []byte {
	return d.body
}
