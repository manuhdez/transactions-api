package service_test

import (
	"fmt"
	"testing"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/infra"
	"github.com/manuhdez/transactions-api/internal/users/test/mocks"
	"github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
	testUser := user.User{}
	hasher := infra.NewBcryptService()

	t.Run("saves user into the repository", func(t *testing.T) {
		repo := mocks.UserMockRepository{Err: nil}
		bus := new(mocks.Bus)
		bus.On("Publish", mock.Anything, mock.Anything).Return(nil)

		srv := service.NewRegisterUserService(repo, hasher, bus)
		got := srv.Register(testUser)
		if got != nil {
			t.Errorf("Register(user): got %v want %v", got, nil)
		}
	})

	t.Run("returns error if cannot save user", func(t *testing.T) {
		want := fmt.Errorf("cannot save user")
		repo := mocks.UserMockRepository{Err: want}
		bus := new(mocks.Bus)
		bus.On("Publish", mock.Anything, mock.Anything).Return(nil)

		srv := service.NewRegisterUserService(repo, hasher, bus)
		got := srv.Register(testUser)
		if got != want {
			t.Errorf("Register(user): got %v want %v", got, nil)
		}
	})
}
