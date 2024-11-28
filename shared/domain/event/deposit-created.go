package event

const depositCreatedType = "event.transactions.deposit_created"

type DepositCreated struct {
	UserId    string  `json:"userId"`
	AccountId string  `json:"accountId"`
	Amount    float32 `json:"amount"`
}

// Type returns the routing key for the deposit created event
func (d DepositCreated) Type() string {
	return depositCreatedType
}

func NewDepositCreated(userId, accountId string, amount float32) *DepositCreated {
	return &DepositCreated{
		UserId:    userId,
		AccountId: accountId,
		Amount:    amount,
	}
}
