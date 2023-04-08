package request_test

import (
	"fmt"
	"testing"

	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/request"
)

type testRequest struct {
	request.Base
}

func (r *testRequest) Validate() {
	r.Errors = append(r.Errors, fmt.Errorf("error #1"))
	r.Errors = append(r.Errors, fmt.Errorf("error #2"))
}

func TestRequest_PrintErrors(t *testing.T) {
	req := testRequest{}
	req.Validate()

	got := req.ErrorResponse()
	want := `{"errors": ["error #1", "error #2"]}`

	if got != want {
		t.Errorf("PrintErrors(errors): got `%s` want `%s`", got, want)
	}
}
