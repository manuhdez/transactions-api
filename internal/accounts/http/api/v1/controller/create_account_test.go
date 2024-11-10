package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/request"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
)

type CreateAccountTestSuite struct {
	suite.Suite
	Account    request.CreateAccount
	Request    *http.Request
	BadRequest *http.Request
	Service    service.CreateService
	Controller CreateAccount
}

func (s *CreateAccountTestSuite) SetupTest() {
	body := bytes.NewBufferString(`{"id": "123", "balance": 100.0}`)
	req := httptest.NewRequest(http.MethodPost, "/accounts", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	s.Request = req
	s.BadRequest = httptest.NewRequest(http.MethodPost, "/accounts", nil)

	repository := new(mocks.AccountMockRepository)
	repository.On("Create", mock.Anything).Return(nil)

	bus := new(mocks.EventBus)
	bus.On("Publish", mock.Anything, mock.Anything).Return(nil)

	s.Service = service.NewCreateService(repository, bus)
	s.Controller = NewCreateAccount(s.Service)
}

func (s *CreateAccountTestSuite) TestCreateAccountWithValidBody() {
	w := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(s.Request, w)

	err := s.Controller.Handle(ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusCreated, w.Code)
}

func TestCreateAccountController(t *testing.T) {
	suite.Run(t, new(CreateAccountTestSuite))
}
