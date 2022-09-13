package event

type AccountCreated struct{}

var AccountCreatedType Type = "event.accounts.account_created"

func (a AccountCreated) Type() Type {
	return AccountCreatedType
}

func (a AccountCreated) Body() []byte {
	return nil
}
