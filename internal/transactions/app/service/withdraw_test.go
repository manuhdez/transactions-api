package service

import (
	"context"
	"errors"
	"testing"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WithdrawTestSuite struct {
	suite.Suite
	repository *mocks.TransactionMockRepository
	Bus        *mocks.EventBus
	service    Withdraw
}

func (s *WithdrawTestSuite) SetupTest() {
	s.repository = new(mocks.TransactionMockRepository)
	s.Bus = new(mocks.EventBus)
	s.service = NewWithdrawService(s.repository, s.Bus)
}

func (s *WithdrawTestSuite) TestShouldCreateNewTransaction() {
	s.repository.On("Withdraw", context.TODO(), mock.Anything).Return(nil)
	withdraw := transaction.NewTransaction(transaction.Withdrawal, "1", 125.5, "EUR")

	res := s.service.Invoke(context.Background(), withdraw)
	assert.NoError(s.T(), res)
}

func (s *WithdrawTestSuite) TestCreateWithdrawError() {
	expected := errors.New("could not create the withdraw")
	s.repository.On("Withdraw", context.TODO(), mock.Anything).Return(expected)
	withdraw := transaction.NewTransaction(transaction.Withdrawal, "23", 33253, "EUR")

	res := s.service.Invoke(context.Background(), withdraw)
	if assert.Error(s.T(), res) {
		assert.Equal(s.T(), expected, res)
	}
}

func TestWithdrawService(t *testing.T) {
	suite.Run(t, new(WithdrawTestSuite))
}
