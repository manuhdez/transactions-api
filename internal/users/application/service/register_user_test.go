package service_test

import (
	"fmt"
	"testing"

	"github.com/manuhdez/transactions-api/internal/users/application/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/infra"
	"github.com/manuhdez/transactions-api/internal/users/test/mocks"
)

func TestRegisterUser(t *testing.T) {

	testUser := user.User{}
	hasher := infra.NewBcryptService()

	t.Run("saves user into the repository", func(t *testing.T) {
		repo := mocks.UserMockRepository{Err: nil}
		srv := service.NewRegisterUserService(repo, hasher)
		got := srv.Register(testUser)

		if got != nil {
			t.Errorf("Register(user): got %v want %v", got, nil)
		}
	})

	t.Run("returns error if cannot save user", func(t *testing.T) {
		want := fmt.Errorf("cannot save user")
		repo := mocks.UserMockRepository{Err: want}
		srv := service.NewRegisterUserService(repo, hasher)

		got := srv.Register(testUser)
		if got != want {
			t.Errorf("Register(user): got %v want %v", got, nil)
		}
	})
}
