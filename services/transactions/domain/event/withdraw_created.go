package event

import (
	"encoding/json"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

var WithdrawCreatedType Type = "event.transactions.withdraw_created"

type WithdrawCreated struct {
	body []byte
}

type WithdrawCreatedBody struct {
	Type    string  `json:"type"`
	Account string  `json:"account"`
	Amount  float32 `json:"amount"`
}

func NewWithdrawCreated(trx transaction.Transaction) WithdrawCreated {
	body := WithdrawCreatedBody{string(WithdrawCreatedType), trx.AccountId, trx.Amount}
	b, _ := json.Marshal(body)
	return WithdrawCreated{b}
}

func (w WithdrawCreated) Type() Type {
	return WithdrawCreatedType
}

func (w WithdrawCreated) Body() []byte {
	return w.body
}
