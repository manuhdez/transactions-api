package controllers

import (
	"encoding/json"
	"io/ioutil"
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
	"github.com/stretchr/testify/require"
)

func TestFindAllController(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/accounts", nil)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = req

	t.Run("returns a list of accounts", func(t *testing.T) {
		accounts := []account.Account{
			account.New("asd", 0),
			account.New("qwe", 23),
			account.New("zxc", 12.5),
		}

		repo := new(mocks.AccountMockRepository)
		repo.On("FindAll", mock.Anything).Return(accounts, nil)

		service := service.NewFindAllService(repo)
		handler := FindAllController(service)
		handler(ctx)

		response := w.Result()
		assert.Equal(t, http.StatusOK, response.StatusCode)

		defer response.Body.Close()

		// Transform account models into json account models and marshal them to get a json string
		acc := infra.NewJsonAccountList(accounts)
		expected, err := json.Marshal(acc)
		require.NoError(t, err)

		body, err := ioutil.ReadAll(response.Body)
		require.NoError(t, err)
		assert.JSONEq(t, string(expected), string(body))
	})
}
