package request

import "fmt"

type Login struct {
	Base
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (l *Login) Validate() {
	if l.Email == "" {
		l.Errors = append(l.Errors, fmt.Errorf("field `email` is required"))
	}
	if l.Password == "" {
		l.Errors = append(l.Errors, fmt.Errorf("field `password` is required"))
	}
}
