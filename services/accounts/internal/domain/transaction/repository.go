package transaction

import "context"

type Repository interface {
	Deposit(ctx context.Context, t Transaction) error
	Withdraw(ctx context.Context, t Transaction) error
	Transfer(ctx context.Context, t Transfer) error
	All(ctx context.Context, userId string) ([]Transaction, error)
	FindByAccount(ctx context.Context, id string) ([]Transaction, error)
}
