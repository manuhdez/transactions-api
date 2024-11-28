package event

import (
	"encoding/json"
	"log"
)

var DepositCreatedType Type = "event.accounts.deposit_created"

type DepositCreated struct {
	body []byte
}

type DepositCreatedBody struct {
	Type    string  `json:"type"`
	Account string  `json:"account"`
	Amount  float32 `json:"amount"`
}

func NewDepositCreated(id string, account string, amount float32) DepositCreated {
	eventType := string(DepositCreatedType)
	body, _ := json.Marshal(DepositCreatedBody{eventType, account, amount})
	return DepositCreated{body}
}

func NewDepositCreatedBody(data []byte) (DepositCreatedBody, error) {
	var body DepositCreatedBody
	err := json.Unmarshal(data, &body)
	if err != nil {
		log.Printf("Error parsing created deposit event: %e", err)
		return DepositCreatedBody{}, err
	}

	return body, nil
}

func (d DepositCreated) Type() Type {
	return DepositCreatedType
}

func (d DepositCreated) Body() []byte {
	return d.body
}
