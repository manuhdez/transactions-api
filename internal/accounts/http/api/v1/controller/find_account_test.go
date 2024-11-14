package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
)

type testSuite struct {
	suite.Suite
	w          *httptest.ResponseRecorder
	req        *http.Request
	ctx        echo.Context
	repository *mocks.AccountMockRepository
	controller FindAccount
}

func (s *testSuite) SetupTest() {
	s.w = httptest.NewRecorder()
	s.req = httptest.NewRequest(http.MethodGet, "/accounts/1", nil)

	e := echo.New()
	s.ctx = e.NewContext(s.req, s.w)
	s.repository = new(mocks.AccountMockRepository)
	s.controller = NewFindAccountController(service.NewFindAccountService(s.repository))
}

func (s *testSuite) TestWithExistingAccount() {
	expected := account.New("123", 33, "EUR")
	s.repository.On("Find", mock.Anything, mock.Anything).Return(expected, nil)
	err := s.controller.Handle(s.ctx)
	assert.NoError(s.T(), err)

	result := s.w.Body.String()
	assert.Equal(s.T(), http.StatusOK, s.w.Code)
	assert.JSONEq(s.T(), infra.JsonStringFromAccount(expected), result)
}

func (s *testSuite) TestWithAccountNotFound() {
	s.repository.On("Find", mock.Anything, mock.Anything).Return(account.Account{}, nil)
	err := s.controller.Handle(s.ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusNotFound, s.w.Code)
}

func TestFindAccountController(t *testing.T) {
	suite.Run(t, new(testSuite))
}
