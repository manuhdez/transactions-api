package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/infra/db"
	"github.com/manuhdez/transactions-api/internal/accounts/test/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup(body io.Reader) (*httptest.ResponseRecorder, echo.Context, *mocks.TransactionRepository) {
	recorder := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/transactions/123", body)
	ctx := echo.New().NewContext(req, recorder)
	repo := new(mocks.TransactionRepository)

	return recorder, ctx, repo
}

func TestFindAccountTransactions(t *testing.T) {
	t.Run("returns a status ok if the request does not fail", func(t *testing.T) {
		rec, ctx, repo := setup(nil)
		repo.On("FindByAccount", mock.Anything, mock.Anything).Return([]transaction.Transaction{}, nil)

		findAccountsController := NewFindAccountTransactions(repo)
		findAccountsController.Handle(ctx)

		want := 200
		got := rec.Code
		assert.EqualValues(t, want, got)
	})

	t.Run("returns an empty list if there are no transactions", func(t *testing.T) {
		rec, ctx, repo := setup(nil)
		repo.On("FindByAccount", mock.Anything, mock.Anything).Return([]transaction.Transaction{}, nil)

		findAccountsController := NewFindAccountTransactions(repo)
		findAccountsController.Handle(ctx)

		got := getJsonBody(t, rec.Body)
		want := response{[]db.JsonTransaction{}}
		assert.Equal(t, want, got)
	})

	t.Run("returns a list of transactions for the given account", func(t *testing.T) {
		rec, ctx, repo := setup(nil)
		repo.On("FindByAccount", mock.Anything, mock.Anything).Return([]transaction.Transaction{
			{AccountId: "1", Type: transaction.Withdrawal, Amount: 112.5},
		}, nil)

		findAccountsController := NewFindAccountTransactions(repo)
		findAccountsController.Handle(ctx)

		got := getJsonBody(t, rec.Body)
		want := response{[]db.JsonTransaction{
			{Account: "1", Type: string(transaction.Withdrawal), Amount: 112.5},
		}}
		assert.Equal(t, want, got)
	})
}

func getJsonBody(t *testing.T, body io.Reader) response {
	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("io.ReadAll(body) err = %e", err)
	}

	var res = response{[]db.JsonTransaction{}}
	err = json.Unmarshal(data, &res)
	if err != nil {
		t.Fatalf("json.Unmarshal(body) err = %e", err)
	}

	return res
}
