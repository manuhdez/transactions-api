package service

import (
	"errors"
	"testing"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CreateServiceTestSuite struct {
	suite.Suite
	Repository *mocks.AccountMockRepository
	Account    account.Account
}

func (suite *CreateServiceTestSuite) SetupTest() {
	suite.Repository = new(mocks.AccountMockRepository)
	suite.Account = account.New("123", 0)
}

func (suite *CreateServiceTestSuite) TestShouldCreateAccount() {
	suite.Repository.On("Create", suite.Account).Return(nil)

	service := NewCreateService(suite.Repository)
	err := service.Create(suite.Account)

	assert.Equal(suite.T(), nil, err)
	suite.Repository.AssertExpectations(suite.T())
}

func (suite *CreateServiceTestSuite) TestShouldReturnError() {
	expected := errors.New("error creating account")
	suite.Repository.On("Create", suite.Account).Return(expected)

	service := NewCreateService(suite.Repository)
	err := service.Create(suite.Account)

	assert.Equal(suite.T(), expected, err)
	suite.Repository.AssertExpectations(suite.T())
}

func TestCreateService(t *testing.T) {
	suite.Run(t, new(CreateServiceTestSuite))
}
