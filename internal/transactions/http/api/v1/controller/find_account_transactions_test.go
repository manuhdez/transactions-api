package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestFindAccountTransactions(t *testing.T) {
	findAccountsController := NewFindAccountTransactions()

	t.Run("returns a status ok if the request does not fail", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		ctx.Request = httptest.NewRequest(http.MethodGet, "/transactions/123", nil)
		findAccountsController.Handle(ctx)

		want := 200
		got := recorder.Code
		assert.Equal(t, want, got)
	})

	t.Run("returns an empty list if there are no transactions", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(recorder)
		ctx.Request = httptest.NewRequest(http.MethodGet, "/transactions/789", nil)
		findAccountsController.Handle(ctx)

		body, err := io.ReadAll(recorder.Body)
		if err != nil {
			log.Fatalf("io.ReadAll(recorder.Body) err = %e", err)
		}

		var got response
		err = json.Unmarshal(body, &got)
		if err != nil {
			t.Fatalf("json.Unmarshal(body) err = %e", err)
		}

		want := response{[]jsonTransaction{}}
		assert.Equal(t, want, got)
	})
}
