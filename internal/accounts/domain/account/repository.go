package account

import "context"

type Repository interface {
	Create(account Account) error
	Find(ctx context.Context, id string) (Account, error)
	GetByUserId(ctx context.Context, userId string) ([]Account, error)
	Delete(ctx context.Context, id string) error
	UpdateBalance(ctx context.Context, id string, balance float32) error
}
