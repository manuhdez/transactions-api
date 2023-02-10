package account

type Account struct {
	Id string
}

func NewAccount(id string) Account {
	return Account{id}
}
