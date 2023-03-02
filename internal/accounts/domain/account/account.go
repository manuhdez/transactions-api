package account

type Account struct {
	id       string
	balance  float32
	currency string
}

func New(id string, balance float32, currency string) Account {
	return Account{id, balance, currency}
}

func (a Account) Id() string {
	return a.id
}

func (a Account) Balance() float32 {
	return a.balance
}

func (a Account) Currency() string {
	return a.currency
}
