package account

import "github.com/manuhdez/transactions-api/shared/domain"

type Account struct {
	id       domain.ID
	UserId   domain.ID
	balance  float32
	currency string
}

func New(id string, balance float32, currency string) Account {
	return Account{
		id:       domain.NewID(id),
		balance:  balance,
		currency: currency,
	}
}

func NewWithUserID(id string, userId domain.ID, balance float32, currency string) Account {
	acc := New(id, balance, currency)
	acc.UserId = userId
	return acc
}

func (a Account) Id() string {
	return a.id.String()
}

func (a Account) Balance() float32 {
	return a.balance
}

func (a Account) Currency() string {
	return a.currency
}
