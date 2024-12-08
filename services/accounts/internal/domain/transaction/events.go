package transaction

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
)

type AccountCreated struct {
	body []byte
}

type AccountCreatedBody struct {
	Type     string  `json:"type"`
	Id       string  `json:"id"`
	UserId   string  `json:"user_id"`
	Balance  float32 `json:"balance"`
	Currency string  `json:"currency"`
}

func NewAccountCreated(acc account.Account) AccountCreated {
	body := AccountCreatedBody{
		Type:     string(event.AccountCreatedType),
		Id:       acc.Id(),
		UserId:   acc.UserId.String(),
		Balance:  acc.Balance(),
		Currency: acc.Currency(),
	}
	b, _ := json.Marshal(body)
	return AccountCreated{b}
}

func (a AccountCreated) Type() event.Type {
	return event.AccountCreatedType
}

func (a AccountCreated) Body() []byte {
	return a.body
}

type DepositCreated struct {
	body []byte
}

type DepositCreatedBody struct {
	Type    string  `json:"type"`
	Account string  `json:"account"`
	Amount  float32 `json:"amount"`
}

func NewDepositCreated(trx Transaction) DepositCreated {
	body := DepositCreatedBody{string(event.DepositCreatedType), trx.AccountId, trx.Amount}
	b, _ := json.Marshal(body)
	return DepositCreated{b}
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

func (d DepositCreated) Type() event.Type {
	return event.DepositCreatedType
}

func (d DepositCreated) Body() []byte {
	return d.body
}

type WithdrawCreated struct {
	body []byte
}

type WithdrawCreatedBody struct {
	Type    string  `json:"type"`
	Account string  `json:"account"`
	Amount  float32 `json:"amount"`
}

func NewWithdrawCreated(trx Transaction) WithdrawCreated {
	body := WithdrawCreatedBody{string(event.WithdrawCreatedType), trx.AccountId, trx.Amount}
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

func (evt WithdrawCreated) Type() event.Type {
	return event.WithdrawCreatedType
}

func (evt WithdrawCreated) Body() []byte {
	return evt.body
}
