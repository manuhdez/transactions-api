package account

type Account struct {
	Id     string
	UserId string
}

func NewAccount(id, userId string) Account {
	return Account{
		Id:     id,
		UserId: userId,
	}
}
