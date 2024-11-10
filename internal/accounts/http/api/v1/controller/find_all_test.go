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

	t.Run("returns a list of accounts", func(t *testing.T) {
		accounts := []account.Account{
			account.New("asd", 0, "EUR"),
			account.New("qwe", 23, "GBP"),
			account.New("zxc", 12.5, "USD"),
		}

		repo := new(mocks.AccountMockRepository)
		repo.On("FindAll", mock.Anything).Return(accounts, nil)

		s := service.NewFindAllService(repo)
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
		acc := infra.NewJsonAccountList(accounts)
		expected, err := json.Marshal(acc)
		require.NoError(t, err)

		body, err := io.ReadAll(response.Body)
		require.NoError(t, err)
		assert.JSONEq(t, string(expected), string(body))
	})
}
