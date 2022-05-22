package infra

import (
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
)

type AccountJson struct {
	Id      string  `json:"id"`
	Balance float32 `json:"balance"`
}

func NewJsonAccount(a account.Account) AccountJson {
	return AccountJson{Id: a.Id(), Balance: a.Balance()}
}

func NewJsonAccountList(list []account.Account) []AccountJson {
	var accounts []AccountJson
	for _, item := range list {
		accounts = append(accounts, NewJsonAccount(item))
	}
	return accounts
}
