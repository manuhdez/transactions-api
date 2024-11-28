package event

import (
	"encoding/json"
	"fmt"
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
