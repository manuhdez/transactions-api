package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type testSuite struct {
	suite.Suite
	w          *httptest.ResponseRecorder
	req        *http.Request
	ctx        *gin.Context
	repository *mocks.AccountMockRepository
	controller FindAccountController
}

func (s *testSuite) SetupTest() {
	s.w = httptest.NewRecorder()
	s.req = httptest.NewRequest(http.MethodGet, "/accounts/1", nil)

	ctx, _ := gin.CreateTestContext(s.w)
	ctx.Request = s.req
	s.ctx = ctx
	s.repository = new(mocks.AccountMockRepository)
	s.controller = NewFindAccountController(service.NewFindAccountService(s.repository))
}

func (s *testSuite) TestWithExistingAccount() {
	expected := account.New("123", 33)
	s.repository.On("Find", mock.Anything, mock.Anything).Return(expected, nil)
	s.controller.Handle(s.ctx)

	result := s.w.Body.String()
	assert.Equal(s.T(), http.StatusOK, s.w.Code)
	assert.Equal(s.T(), infra.JsonStringFromAccount(expected), result)
}

func (s *testSuite) TestWithAccountNotFound() {
	s.repository.On("Find", mock.Anything, mock.Anything).Return(account.Account{}, nil)
	s.controller.Handle(s.ctx)
	assert.Equal(s.T(), http.StatusNotFound, s.w.Code)
}

func TestFindAccountController(t *testing.T) {
	suite.Run(t, new(testSuite))
}
