package event

const withdrawCreatedType = "event.transactions.withdraw_created"

type WithdrawCreated struct {
	UserId    string  `json:"userId"`
	AccountId string  `json:"accountId"`
	Amount    float32 `json:"amount"`
}

// Type returns the routing key for the withdraw created event
func (w WithdrawCreated) Type() string {
	return withdrawCreatedType
}

func NewWithdrawCreated(userId, accountId string, amount float32) *WithdrawCreated {
	return &WithdrawCreated{
		UserId:    userId,
		AccountId: accountId,
		Amount:    amount,
	}
}
