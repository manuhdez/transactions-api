package account

import "context"

type Repository interface {
	Save(ctx context.Context, account Account) error
	FindById(ctx context.Context, id string) (Account, error)
}
