package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/manuhdez/transactions-api/internal/accounts/http/api/v1/request"
	sharedhttp "github.com/manuhdez/transactions-api/shared/infra/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
)

type CreateAccountTestSuite struct {
	suite.Suite

	srv     *echo.Echo
	w       *httptest.ResponseRecorder
	repo    *mocks.AccountMockRepository
	bus     *mocks.EventBus
	creator service.CreateService
	ctrl    CreateAccount
}

func (s *CreateAccountTestSuite) SetupTest() {
	s.srv = echo.New()
	s.srv.Validator = sharedhttp.NewRequestValidator()
	s.w = httptest.NewRecorder()

	s.repo = new(mocks.AccountMockRepository)
	s.bus = new(mocks.EventBus)

	s.creator = service.NewCreateService(s.repo, s.bus)
	s.ctrl = NewCreateAccount(s.creator)
}

func (s *CreateAccountTestSuite) TestCreateAccountWithValidBody() {
	body, err := json.Marshal(request.CreateAccount{Id: "123", UserId: "321", Currency: ""})
	require.NoError(s.T(), err)
	req := httptest.NewRequest(http.MethodPost, "/accounts", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	ctx := s.srv.NewContext(req, s.w)
	ctx.Set("userId", "321")

	s.repo.On("Create", mock.Anything, mock.Anything).Return(nil)
	s.bus.On("Publish", mock.Anything, mock.Anything).Return(nil)

	err = s.ctrl.Handle(ctx)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), http.StatusCreated, s.w.Code)
}

func TestCreateAccountController(t *testing.T) {
	suite.Run(t, new(CreateAccountTestSuite))
}
