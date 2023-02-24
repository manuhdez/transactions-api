package controller

import (
	"encoding/json"
	"github.com/manuhdez/transactions-api/internal/transactions/infra"
	"github.com/manuhdez/transactions-api/internal/transactions/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
)

func setup(body io.Reader) (*httptest.ResponseRecorder, *gin.Context, *mocks.TransactionMockRepository) {
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodGet, "/transactions/123", body)
	repo := new(mocks.TransactionMockRepository)
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
		want := response{[]infra.JsonTransaction{}}
		assert.Equal(t, want, got)
	})

	t.Run("returns a list of transactions for the given account", func(t *testing.T) {
		rec, ctx, repo := setup(nil)
		repo.On("FindByAccount", mock.Anything, mock.Anything).Return([]transaction.Transaction{
			{AccountId: "1", Type: transaction.Withdrawal, Amount: 112.5, Currency: "EUR"},
		}, nil)

		findAccountsController := NewFindAccountTransactions(repo)
		findAccountsController.Handle(ctx)

		got := getJsonBody(t, rec.Body)
		want := response{[]infra.JsonTransaction{
			{Account: "1", Type: string(transaction.Withdrawal), Amount: 112.5, Currency: "EUR"},
		}}
		assert.Equal(t, want, got)
	})
}

func getJsonBody(t *testing.T, body io.Reader) response {
	data, err := io.ReadAll(body)
	if err != nil {
		t.Fatalf("io.ReadAll(body) err = %e", err)
	}

	var res = response{[]infra.JsonTransaction{}}
	err = json.Unmarshal(data, &res)
	if err != nil {
		t.Fatalf("json.Unmarshal(body) err = %e", err)
	}

	return res
}
