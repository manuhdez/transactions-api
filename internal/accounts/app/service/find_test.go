package service

import (
	"context"
	"errors"
	"testing"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	repository *mocks.AccountMockRepository
	service    FindAccountService
}

func (s *testSuite) SetupTest() {
	s.repository = new(mocks.AccountMockRepository)
	s.service = NewFindAccountService(s.repository)
}

func (s *testSuite) TestWithMatchingAccount() {
	expected := account.New("123", 32, "EUR")
	s.repository.On("Find", mock.Anything, mock.Anything).Return(expected, nil)

	result, err := s.service.Find(context.Background(), "123")
	require.NoError(s.T(), err)
	assert.Equal(s.T(), expected, result)
}

func (s *testSuite) TestAccountNotFoundThrowsError() {
	expected := errors.New("account not found")
	s.repository.On("Find", mock.Anything, mock.Anything).Return(account.Account{}, expected)

	res, err := s.service.Find(context.Background(), "123")
	assert.Error(s.T(), err)
	assert.ErrorContains(s.T(), err, "not found")
	assert.Equal(s.T(), account.Account{}, res)
}

func (s *testSuite) TestWithError() {
	s.repository.On("Find", mock.Anything, mock.Anything).Return(account.Account{}, assert.AnError)

	res, err := s.service.Find(context.Background(), "123")
	assert.Error(s.T(), err)
	assert.Equal(s.T(), assert.AnError, err)
	assert.Equal(s.T(), account.Account{}, res)
}

func TestFindAccountService(t *testing.T) {
	suite.Run(t, new(testSuite))
}
