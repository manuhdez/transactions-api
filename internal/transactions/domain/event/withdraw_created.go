package event

import "encoding/json"

var WithdrawCreatedType Type = "event.transactions.withdraw_created"

type WithdrawCreated struct {
	body []byte
}

type WithdrawCreatedBody struct {
	Type    string  `json:"type"`
	Account string  `json:"account"`
	Amount  float32 `json:"amount"`
}

func NewWithdrawCreated(account string, amount float32) WithdrawCreated {
	body := WithdrawCreatedBody{string(WithdrawCreatedType), account, amount}
	b, _ := json.Marshal(body)
	return WithdrawCreated{b}
}

func (w WithdrawCreated) Type() Type {
	return WithdrawCreatedType
}

func (w WithdrawCreated) Body() []byte {
	return w.body
}
