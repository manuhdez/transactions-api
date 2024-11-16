package transaction

type Type string

const (
	Deposit      Type = "deposit"
	Withdrawal   Type = "withdrawal"
	TransferType Type = "transfer"
)

type Transaction struct {
	Type      Type
	AccountId string
	UserId    string
	Amount    float32
}

func NewTransaction(transactionType Type, accountId, userId string, amount float32) Transaction {
	return Transaction{
		Type:      transactionType,
		AccountId: accountId,
		UserId:    userId,
		Amount:    amount,
	}
}

type Transfer struct {
	From   string  // origin account
	To     string  // destination account
	Amount float32 // amount to transfer
}
