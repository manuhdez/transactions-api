package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"
)

func TestFindAllController(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/accounts", nil)
	w := httptest.NewRecorder()
	e := echo.New()
	ctx := e.NewContext(req, w)

	t.Run("returns the list of accounts of the user", func(t *testing.T) {
		accounts := []account.Account{
			account.NewWithUserID("asd", "123", 0, "EUR"),
			account.NewWithUserID("qwe", "123", 23, "GBP"),
			account.NewWithUserID("zxc", "123", 12.5, "USD"),
		}

		repo := new(mocks.AccountMockRepository)
		repo.On("GetByUserId", mock.Anything, mock.Anything).Return(accounts, nil)

		ctx.Set("userId", "123")

		s := service.NewFindAccountService(repo)
		err := NewFindAllAccounts(s).Handle(ctx)
		assert.NoError(t, err)

		response := w.Result()
		assert.Equal(t, http.StatusOK, response.StatusCode)

		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(response.Body)

		// Transform account models into json account models and marshal them to get a json string
		expected, err := json.Marshal(findAllAccountsResponse{
			Accounts: infra.NewJsonAccountList(accounts),
		})
		require.NoError(t, err)

		body := w.Body.String()
		assert.JSONEq(t, string(expected), body)
	})
}
