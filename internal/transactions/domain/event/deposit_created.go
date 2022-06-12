package event

type DepositCreated struct{}

var DepositCreatedType Type = "event.transactions.deposit_created"

func (d DepositCreated) Type() Type {
	return DepositCreatedType
}
