package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CreateAccountTestSuite struct {
	suite.Suite
	Account    createAccountRequest
	Request    *http.Request
	BadRequest *http.Request
	Service    service.CreateService
	Controller CreateAccountController
}

func (s *CreateAccountTestSuite) SetupTest() {
	body := bytes.NewBufferString(`{"id": "123", "balance": 100.0}`)
	s.Request = httptest.NewRequest(http.MethodPost, "/accounts", body)
	s.BadRequest = httptest.NewRequest(http.MethodPost, "/accounts", nil)

	repository := new(mocks.AccountMockRepository)
	repository.On("Create", mock.Anything).Return(nil)

	bus := new(mocks.EventBus)
	bus.On("Publish", mock.Anything, mock.Anything).Return(nil)

	s.Service = service.NewCreateService(repository, bus)
	s.Controller = NewCreateAccountController(s.Service)
}

func (s *CreateAccountTestSuite) TestCreateAccountWithValidBody() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = s.Request

	s.Controller.Handle(ctx)

	res := w.Result()
	assert.Equal(s.T(), http.StatusCreated, res.StatusCode)
}

func (s *CreateAccountTestSuite) TestCreateAccountWithBadRequest() {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = s.BadRequest

	s.Controller.Handle(ctx)

	res := w.Result()
	assert.Equal(s.T(), http.StatusBadRequest, res.StatusCode)
}

func TestCreateAccountController(t *testing.T) {
	suite.Run(t, new(CreateAccountTestSuite))
}
