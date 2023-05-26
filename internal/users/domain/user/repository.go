package user

import "context"

//go:generate mockery --case=snake --outpkg=mocks --output="../../test/mocks" --name=Repository

type Repository interface {
	All(ctx context.Context) ([]User, error)
	Save(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
}
