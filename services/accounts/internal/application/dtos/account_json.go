package dtos

import (
	"encoding/json"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
)

type AccountJson struct {
	Id       string  `json:"id"`
	UserId   string  `json:"user_id"`
	Balance  float32 `json:"balance"`
	Currency string  `json:"currency"`
}

func NewJsonAccount(a account.Account) AccountJson {
	return AccountJson{
		Id:       a.Id(),
		UserId:   a.UserId.String(),
		Balance:  a.Balance(),
		Currency: a.Currency(),
	}
}

func NewJsonAccountList(list []account.Account) []AccountJson {
	accounts := make([]AccountJson, len(list))
	for i := range list {
		accounts[i] = NewJsonAccount(list[i])
	}
	return accounts
}

func JsonStringFromAccount(a account.Account) string {
	str, _ := json.Marshal(NewJsonAccount(a))
	return string(str)
}
