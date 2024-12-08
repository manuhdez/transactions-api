package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/dtos"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/application/service"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
)

type findAccountsTestSuite struct {
	suite.Suite
	w          *httptest.ResponseRecorder
	req        *http.Request
	ctx        echo.Context
	repository *mocks.AccountRepository
	controller FindAccount
}

func (s *findAccountsTestSuite) SetupTest() {
	s.w = httptest.NewRecorder()
	s.req = httptest.NewRequest(http.MethodGet, "/accounts/1", nil)

	e := echo.New()
	s.ctx = e.NewContext(s.req, s.w)
	s.repository = new(mocks.AccountRepository)
	s.controller = NewFindAccountController(service.NewFindAccountService(s.repository))
}

func (s *findAccountsTestSuite) TestWithExistingAccount() {
	expected := account.New("123", 33, "EUR")
	s.repository.On("Find", mock.Anything, mock.Anything).Return(expected, nil)
	err := s.controller.Handle(s.ctx)
	assert.NoError(s.T(), err)

	result := s.w.Body.String()
	assert.Equal(s.T(), http.StatusOK, s.w.Code)
	assert.JSONEq(s.T(), dtos.JsonStringFromAccount(expected), result)
}

func (s *findAccountsTestSuite) TestWithAccountNotFound() {
	s.repository.On("Find", mock.Anything, mock.Anything).Return(account.Account{}, nil)
	err := s.controller.Handle(s.ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusNotFound, s.w.Code)
}

func TestFindAccountController(t *testing.T) {
	suite.Run(t, new(testSuite))
}
