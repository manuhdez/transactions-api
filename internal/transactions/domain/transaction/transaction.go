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

func NewTransaction(t Type, acc, user string, amount float32) Transaction {
	return Transaction{Type: t, AccountId: acc, UserId: user, Amount: amount}
}

func NewDeposit(acc, user string, amount float32) Transaction {
	return Transaction{Type: Deposit, AccountId: acc, UserId: user, Amount: amount}
}

func NewWithdraw(acc, user string, amount float32) Transaction {
	return Transaction{Type: Withdrawal, AccountId: acc, UserId: user, Amount: amount}
}

type Transfer struct {
	UserId string  // user performing the transfer
	From   string  // origin account
	To     string  // destination account
	Amount float32 // amount to transfer
}

func NewTransfer(user, from, to string, amount float32) Transfer {
	return Transfer{UserId: user, From: from, To: to, Amount: amount}
}
