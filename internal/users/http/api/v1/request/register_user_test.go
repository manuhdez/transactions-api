package request_test

import (
	"testing"

	"github.com/manuhdez/transactions-api/internal/users/http/api/v1/request"
)

type testCase struct {
	name string
	req  request.RegisterUser
	want int
}

func TestRegisterUser_Validate(t *testing.T) {
	for _, tt := range []testCase{
		{
			name: "with a valid request",
			req:  request.RegisterUser{Id: "1", FirstName: "Name", LastName: "Last", Email: "mail@test.com", Password: "pass"},
			want: 0,
		},
		{
			name: "`Id` field is required",
			req:  request.RegisterUser{FirstName: "Name", LastName: "Last", Email: "mail@test.com", Password: "pass"},
			want: 1,
		},
		{
			name: "`FirstName` field is required",
			req:  request.RegisterUser{Id: "1", LastName: "Last", Email: "mail@test.com", Password: "pass"},
			want: 1,
		},
		{
			name: "`LasName` field is required",
			req:  request.RegisterUser{Id: "2", FirstName: "Name", Email: "mail@test.com", Password: "pass"},
			want: 1,
		},
		{
			name: "`Email` field is required",
			req:  request.RegisterUser{Id: "1", FirstName: "first", LastName: "last", Password: "xyz"},
			want: 1,
		},
		{
			name: "`Email` field has a wrong format",
			req:  request.RegisterUser{Id: "1", FirstName: "first", LastName: "last", Email: "email", Password: "123"},
			want: 1,
		},
		{
			name: "with an invalid `Email` format",
			req:  request.RegisterUser{Id: "1", FirstName: "first", LastName: "last", Email: "mail", Password: "pas"},
			want: 1,
		},
		{
			name: "`Password` field is required",
			req:  request.RegisterUser{Id: "1", FirstName: "Name", LastName: "Last", Email: "mail@test.com"},
			want: 1,
		},
		{
			name: "with two missing fields",
			req:  request.RegisterUser{LastName: "Last", Email: "mail@test.com", Password: "pas"},
			want: 2,
		},
		{
			name: "missing all required fields",
			req:  request.RegisterUser{},
			want: 5,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			errors := tt.req.Validate()

			got := len(errors)
			if got != tt.want {
				t.Errorf("Validate(): got %v want %v", got, tt.want)
			}
		})
	}
}
