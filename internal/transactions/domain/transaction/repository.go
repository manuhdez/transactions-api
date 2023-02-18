package transaction

import "context"

type Repository interface {
	Deposit(ctx context.Context, t Transaction) error
    Withdraw(ctx context.Context, t Transaction) error
	FindAll(ctx context.Context) ([]Transaction, error)
}
