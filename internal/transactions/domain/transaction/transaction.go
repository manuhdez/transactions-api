package transaction

type Type string

const (
	Deposit    Type = "deposit"
	Withdrawal Type = "withdrawal"
)

type Transaction struct {
	Type      Type
	AccountId string
	Amount    float32
	Currency  string
}

func NewTransaction(transactionType Type, accountId string, amount float32, currency string) Transaction {
	return Transaction{
		Type:      transactionType,
		AccountId: accountId,
		Amount:    amount,
		Currency:  currency,
	}
}
