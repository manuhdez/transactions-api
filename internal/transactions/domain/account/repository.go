package account

import "context"

type Repository interface {
    Save(ctx context.Context, account Account) error
}
