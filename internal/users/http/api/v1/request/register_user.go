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

func (r *RegisterUser) Validate() {
	if r.Id == "" {
		r.Errors = append(r.Errors, fmt.Errorf("field `id` is required"))
	}
	if r.FirstName == "" {
		r.Errors = append(r.Errors, fmt.Errorf("field `first_name` is required"))
	}
	if r.LastName == "" {
		r.Errors = append(r.Errors, fmt.Errorf("field `last_name` is required"))
	}
	if r.Email == "" {
		r.Errors = append(r.Errors, fmt.Errorf("field `email` is required"))
	}
	if strings.Contains(r.Email, "@") == false {
		r.Errors = append(r.Errors, fmt.Errorf("field `email` does not match the required format"))
	}
	if r.Password == "" {
		r.Errors = append(r.Errors, fmt.Errorf("field `password` is required"))
	}
}
