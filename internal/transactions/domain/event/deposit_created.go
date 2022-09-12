package event

import "encoding/json"

type DepositCreated struct {
	body []byte
}

type DepositCreatedBody struct {
	Type    string  `json:"type"`
	Account string  `json:"account"`
	Amount  float32 `json:"amount"`
}

func NewDepositCreated(account string, amount float32) DepositCreated {
	body := DepositCreatedBody{string(DepositCreatedType), account, amount}
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
