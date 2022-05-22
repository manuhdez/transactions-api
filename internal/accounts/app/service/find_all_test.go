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

type FindAllSuite struct {
	suite.Suite
	repository *mocks.AccountMockRepository
	service    FindAllService
	ctx        context.Context
}

func (s *FindAllSuite) SetupTest() {
	s.repository = new(mocks.AccountMockRepository)
	s.service = NewFindAllService(s.repository)
	s.ctx = context.Background()
}

func (s *FindAllSuite) TestWithEmptyListOfAccounts() {
	s.repository.On("FindAll", mock.Anything).Return([]account.Account{}, nil)

	result, err := s.service.Find(s.ctx)
	require.NoError(s.T(), err)

	assert.Equal(s.T(), 0, len(result))
}

func (s *FindAllSuite) TestWithAListOfAccounts() {
	expected := []account.Account{
		account.New("1", 10),
		account.New("2", 23.95),
	}
	s.repository.On("FindAll", mock.Anything).Return(expected, nil)

	result, err := s.service.Find(s.ctx)
	require.NoError(s.T(), err)
	assert.Equal(s.T(), expected, result)
}

func (s *FindAllSuite) TestReturnsErrorIfRepositoryFails() {
	expected := errors.New("cannot find accounts")
	s.repository.On("FindAll", mock.Anything).Return([]account.Account{}, expected)

	result, err := s.service.Find(s.ctx)
	assert.Len(s.T(), result, 0)
	assert.ErrorContains(s.T(), err, "cannot find accounts")
}

func TestFindAllService(t *testing.T) {
	suite.Run(t, new(FindAllSuite))
}
