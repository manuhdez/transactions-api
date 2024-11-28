package transaction

import "context"

type Repository interface {
	Deposit(ctx context.Context, t Transaction) error
	Withdraw(ctx context.Context, t Transaction) error
	All(ctx context.Context, userId string) ([]Transaction, error)
	FindByAccount(ctx context.Context, id string) ([]Transaction, error)
}
