package event

import (
	"encoding/json"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
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

var AccountCreatedType Type = "event.accounts.account_created"

func NewAccountCreated(acc account.Account) AccountCreated {
	body := AccountCreatedBody{
		Type:     string(AccountCreatedType),
		Id:       acc.Id(),
		UserId:   acc.UserId.String(),
		Balance:  acc.Balance(),
		Currency: acc.Currency(),
	}
	b, _ := json.Marshal(body)
	return AccountCreated{b}
}

func (a AccountCreated) Type() Type {
	return AccountCreatedType
}

func (a AccountCreated) Body() []byte {
	return a.body
}
