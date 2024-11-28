package request_test

import (
	"fmt"
	"testing"

	"github.com/manuhdez/transactions-api/internal/users/internal/api/http/v1/request"
)

type testRequest struct {
	request.Base
}

func (r *testRequest) Validate() []error {
	var errors []error
	errors = append(errors, fmt.Errorf("error #1"))
	return append(errors, fmt.Errorf("error #2"))
}

func TestRequest_PrintErrors(t *testing.T) {
	req := testRequest{}
	errorsList := req.Validate()

	got := req.ErrorResponse(errorsList)
	want := `{"errors": ["error #1", "error #2"]}`

	if got != want {
		t.Errorf("PrintErrors(errors): got `%s` want `%s`", got, want)
	}
}
