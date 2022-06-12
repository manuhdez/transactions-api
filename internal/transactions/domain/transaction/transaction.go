package transaction

type transactionType string

const (
	Deposit    transactionType = "deposit"
	Withdrawal transactionType = "withdrawal"
)

type Transaction struct {
	Type      transactionType
	AccountId string
	Amount    float32
	Currency  string
}

func NewTransaction(transactionType transactionType, accountId string, amount float32, currency string) Transaction {
	return Transaction{
		Type:      transactionType,
		AccountId: accountId,
		Amount:    amount,
		Currency:  currency,
	}
}
