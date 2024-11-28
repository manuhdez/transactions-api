package event

import (
	"encoding/json"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
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
func NewWithdrawCreatedBody(data []byte) (WithdrawCreatedBody, error) {
	var body WithdrawCreatedBody
	err := json.Unmarshal(data, &body)
	if err != nil {
		fmt.Printf("Could not parse event body: %s", err)
		return WithdrawCreatedBody{}, err
	}

	return body, nil
}

func (evt WithdrawCreated) Type() Type {
	return WithdrawCreatedType
}

func (evt WithdrawCreated) Body() []byte {
	return evt.body
}
