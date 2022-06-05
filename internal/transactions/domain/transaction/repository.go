package transaction

import "context"

type Repository interface {
	Deposit(ctx context.Context, t Transaction) error
}
