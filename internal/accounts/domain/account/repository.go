package account

import "context"

type Repository interface {
	Create(account Account) error
	FindAll(ctx context.Context) ([]Account, error)
	Find(ctx context.Context, id string) (Account, error)
	Delete(ctx context.Context, id string) error
	UpdateBalance(ctx context.Context, id string, balance float32) error
}
