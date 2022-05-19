package account

type Account struct {
	Id      string
	Balance float32
}

func New(id string, balance float32) Account {
	return Account{id, balance}
}
