package account

type Account struct {
	id      string
	balance float32
}

func New(id string, balance float32) Account {
	return Account{id, balance}
}

func (a Account) Id() string {
	return a.id
}

func (a Account) Balance() float32 {
	return a.balance
}
