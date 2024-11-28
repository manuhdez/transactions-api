package account

import "context"

type Repository interface {
	Create(context.Context, Account) error
	Find(context.Context, string) (Account, error)
	GetByUserId(context.Context, string) ([]Account, error)
	Delete(context.Context, string) error
	UpdateBalance(context.Context, string, float32) error
}
