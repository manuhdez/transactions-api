package service

import (
	"context"
	"errors"
	"testing"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateServiceTestSuite struct {
	suite.Suite
	Repository *mocks.AccountMockRepository
	Bus        *mocks.EventBus
	Account    account.Account
}

func (suite *CreateServiceTestSuite) SetupTest() {
	suite.Repository = new(mocks.AccountMockRepository)
	suite.Bus = new(mocks.EventBus)
	suite.Account = account.New("123", 0, "EUR")
}

func (suite *CreateServiceTestSuite) TestShouldCreateAccount() {
	suite.Repository.On("Create", suite.Account).Return(nil)
	suite.Bus.On("Publish", context.Background(), mock.Anything).Return(nil)

	service := NewCreateService(suite.Repository, suite.Bus)
	err := service.Create(suite.Account)

	assert.Equal(suite.T(), nil, err)
	suite.Repository.AssertExpectations(suite.T())
}

func (suite *CreateServiceTestSuite) TestShouldReturnError() {
	expected := errors.New("error creating account")
	suite.Repository.On("Create", suite.Account).Return(expected)
	suite.Bus.On("Publish", context.Background(), mock.Anything).Return(nil)

	service := NewCreateService(suite.Repository, suite.Bus)
	err := service.Create(suite.Account)

	assert.Equal(suite.T(), expected, err)
	suite.Repository.AssertExpectations(suite.T())
}

func TestCreateService(t *testing.T) {
	suite.Run(t, new(CreateServiceTestSuite))
}
