package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatusController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := httptest.NewRequest(http.MethodGet, "/status", nil)

	t.Run("should return status 200", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = r
		StatusController(ctx)

		res := w.Result()
		assert.Equal(t, res.StatusCode, http.StatusOK)
	})

	t.Run("should return the correct json response", func(t *testing.T) {
		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = r
		StatusController(ctx)

		result, err := ioutil.ReadAll(w.Result().Body)
		require.NoError(t, err)

		expected, err := json.Marshal(statusResponse{"ok", "accounts service"})
		require.NoError(t, err)

		assert.Equal(t, expected, result)
	})
}
