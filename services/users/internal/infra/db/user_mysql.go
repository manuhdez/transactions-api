package db

import (
	"github.com/manuhdez/transactions-api/internal/users/internal/domain/user"
)

type UserPostgres struct {
	Id        string `mysql:"id"`
	FirstName string `mysql:"first_name"`
	LastName  string `mysql:"last_name"`
	Email     string `mysql:"email"`
	Password  string `mysql:"password"`
}

func (u UserPostgres) TableName() string {
	return "users"
}

func (u UserPostgres) ToDomainModel() user.User {
	return user.New(u.Id, u.FirstName, u.LastName, u.Email, u.Password)
}
