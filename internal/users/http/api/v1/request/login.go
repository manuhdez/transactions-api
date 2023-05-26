package request

import "fmt"

type Login struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *Login) Validate() []error {
	var errors []error
	if l.Email == "" {
		errors = append(errors, fmt.Errorf("field `email` is required"))
	}
	if l.Password == "" {
		errors = append(errors, fmt.Errorf("field `password` is required"))
	}
	return errors
}
