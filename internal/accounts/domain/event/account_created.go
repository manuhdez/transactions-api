package event

import "encoding/json"

type AccountCreated struct {
	body []byte
}

type AccountCreatedBody struct {
	Type     string  `json:"type"`
	Id       string  `json:"id"`
	Balance  float32 `json:"balance"`
	Currency string  `json:"currency"`
}

var AccountCreatedType Type = "event.accounts.account_created"

func NewAccountCreated(id string, balance float32, currency string) AccountCreated {
	body := AccountCreatedBody{string(AccountCreatedType), id, balance, currency}
	b, _ := json.Marshal(body)
	return AccountCreated{b}
}

func (a AccountCreated) Type() Type {
	return AccountCreatedType
}

func (a AccountCreated) Body() []byte {
	return a.body
}
