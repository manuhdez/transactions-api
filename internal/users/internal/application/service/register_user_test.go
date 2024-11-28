package service_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/users/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/users/internal/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/internal/infra"
	"github.com/manuhdez/transactions-api/internal/users/test/mocks"
)

var (
	errSavingUser = errors.New("error saving user")
)

type registerUserSuite struct {
	repo   *mocks.UserRepository
	hasher *mocks.Hasher
	bus    *mocks.EventBus
}

func setupRegisterUserSuite() registerUserSuite {
	return registerUserSuite{
		repo:   new(mocks.UserRepository),
		hasher: new(mocks.Hasher),
		bus:    new(mocks.EventBus),
	}
}

func (s registerUserSuite) assertMocks(t *testing.T) {
	s.repo.AssertExpectations(t)
	s.bus.AssertExpectations(t)
}

func TestRegisterUser(t *testing.T) {
	testUser := user.User{}
	hasher := infra.NewBcryptService()

	t.Run("saves user into the repository", func(t *testing.T) {
		s := setupRegisterUserSuite()
		defer s.assertMocks(t)

		s.repo.On("Save", mock.Anything, mock.Anything).Return(nil).Once()
		s.hasher.On("Hash", mock.Anything).Return("hashed-password").Once()
		s.bus.On("Publish", mock.Anything, mock.Anything).Return(nil).Once()

		srv := service.NewRegisterUserService(s.repo, hasher, s.bus)
		got := srv.Register(testUser)
		if got != nil {
			t.Errorf("Register(user): got %v want %v", got, nil)
		}
	})

	t.Run("returns error if cannot save user", func(t *testing.T) {
		s := setupRegisterUserSuite()
		defer s.assertMocks(t)

		s.repo.On("Save", mock.Anything, mock.Anything).Return(errSavingUser)

		srv := service.NewRegisterUserService(s.repo, hasher, s.bus)
		if got := srv.Register(testUser); !errors.Is(got, errSavingUser) {
			t.Errorf("Register(user): got %v want %v", got, nil)
		}
	})
}
