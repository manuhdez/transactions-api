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

func (a Account) Id() string {
	return a.id.String()
}

func (a Account) Balance() float32 {
	return a.balance
}

func (a Account) Currency() string {
	return a.currency
}
