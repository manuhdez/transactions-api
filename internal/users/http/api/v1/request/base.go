package request

import (
	"fmt"
	"strings"
)

type Request interface {
	Validate() []error
	ErrorResponse() string
}

// Base - is an abstract implementation of a Request object
// The Validate method needs to be overridden by the actual Request implementation
// The ErrorResponse method has a default implementation that returns a valid json response
type Base struct {
}

// Validate - Checks the validity of the request fields and stores the errors
// Returns a bool that represent if the data is valid or not
func (r *Base) Validate() []error {
	return nil
}

// ErrorResponse - Generates a string representation of a list of errors in json format
func (r *Base) ErrorResponse(errorList []error) string {
	var str []string
	for _, err := range errorList {
		str = append(str, fmt.Sprintf(`"%s"`, err.Error()))
	}
	return fmt.Sprintf(`{"errors": [%s]}`, strings.Join(str, ", "))
}
