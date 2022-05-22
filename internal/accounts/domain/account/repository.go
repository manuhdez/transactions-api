package account

import "context"

type Repository interface {
	Create(account Account) error
	FindAll(ctx context.Context) ([]Account, error)
}
