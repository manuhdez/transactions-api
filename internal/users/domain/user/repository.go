package user

import "context"

type Repository interface {
	Save(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
}
