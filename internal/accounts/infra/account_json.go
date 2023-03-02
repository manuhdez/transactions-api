package infra

import (
	"encoding/json"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type AccountJson struct {
	Id       string  `json:"id"`
	Balance  float32 `json:"balance"`
	Currency string  `json:"currency"`
}

func NewJsonAccount(a account.Account) AccountJson {
	return AccountJson{Id: a.Id(), Balance: a.Balance(), Currency: a.Currency()}
}

func NewJsonAccountList(list []account.Account) []AccountJson {
	var accounts []AccountJson
	for _, item := range list {
		accounts = append(accounts, NewJsonAccount(item))
	}
	return accounts
}

func JsonStringFromAccount(a account.Account) string {
	str, _ := json.Marshal(NewJsonAccount(a))
	return string(str)
}
