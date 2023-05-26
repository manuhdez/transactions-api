package request

import (
	"fmt"
	"strings"
)

type RegisterUser struct {
	Base
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

func (r *RegisterUser) Validate() []error {
	var errors []error
	if r.Id == "" {
		errors = append(errors, fmt.Errorf("field `id` is required"))
	}

	if r.FirstName == "" {
		errors = append(errors, fmt.Errorf("field `first_name` is required"))
	}

	if r.LastName == "" {
		errors = append(errors, fmt.Errorf("field `last_name` is required"))
	}

	if r.Email == "" {
		errors = append(errors, fmt.Errorf("field `email` is required"))
	} else if strings.Contains(r.Email, "@") == false {
		errors = append(errors, fmt.Errorf("field `email` does not match the required format"))
	}

	if r.Password == "" {
		errors = append(errors, fmt.Errorf("field `password` is required"))
	}

	return errors
}
