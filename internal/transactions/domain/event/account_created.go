package event

import (
	"encoding/json"
	"log"
)

const AccountCreatedType Type = "event.accounts.account_created"

type AccountCreated struct {
	body []byte
}

func NewAccountCreated(data []byte) (AccountCreated, error) {
	body := AccountCreatedBody{string(AccountCreatedType)}
	b, err := json.Marshal(body)
	if err != nil {
		return AccountCreated{}, err
	}

	return AccountCreated{b}, nil
}

func (a AccountCreated) Type() Type {
	return AccountCreatedType
}

func (a AccountCreated) Body() []byte {
	return a.body
}

type AccountCreatedBody struct {
	Id string `json:"id"`
}

func NewAccountCreatedBody(data []byte) (AccountCreatedBody, error) {
	var body AccountCreatedBody
	err := json.Unmarshal(data, &body)
	if err != nil {
		log.Printf("Error parsing created account event: %e", err)
		return AccountCreatedBody{}, err
	}
	return body, nil
}
