package user

import "context"

type Repository interface {
	All(ctx context.Context) ([]User, error)
	Save(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
}
