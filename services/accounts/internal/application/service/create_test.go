package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"

	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
)

type CreateServiceTestSuite struct {
	suite.Suite
	Repository *mocks.AccountMockRepository
	Bus        *mocks.EventBus
	Account    account.Account
	ctx        context.Context
}

func (s *CreateServiceTestSuite) SetupTest() {
	s.Repository = new(mocks.AccountMockRepository)
	s.Bus = new(mocks.EventBus)
	s.Account = account.NewWithUserID("333", "123", 0, "EUR")
	s.ctx = context.Background()
}

func (s *CreateServiceTestSuite) TestShouldCreateAccount() {
	s.Repository.On("Create", mock.Anything, s.Account).Return(nil)
	s.Bus.On("Publish", context.Background(), mock.Anything).Return(nil)

	service := NewCreateService(s.Repository, s.Bus)
	err := service.Create(s.ctx, s.Account)

	assert.Equal(s.T(), nil, err)
	s.Repository.AssertExpectations(s.T())
}

func (s *CreateServiceTestSuite) TestShouldReturnError() {
	expected := errors.New("error creating account")
	s.Repository.On("Create", mock.Anything, s.Account).Return(expected)
	s.Bus.On("Publish", context.Background(), mock.Anything).Return(nil)

	service := NewCreateService(s.Repository, s.Bus)
	err := service.Create(s.ctx, s.Account)

	assert.ErrorIs(s.T(), err, expected)
	s.Repository.AssertExpectations(s.T())
}

func TestCreateService(t *testing.T) {
	suite.Run(t, new(CreateServiceTestSuite))
}
