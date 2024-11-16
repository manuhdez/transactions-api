package infra

import "github.com/manuhdez/transactions-api/internal/users/domain/user"

type UserMysql struct {
	Id        string `mysql:"id"`
	FirstName string `mysql:"first_name"`
	LastName  string `mysql:"last_name"`
	Email     string `mysql:"email"`
	Password  string `mysql:"password"`
}

func (u UserMysql) TableName() string {
	return "users"
}

func (u UserMysql) ToDomainModel() user.User {
	return user.New(u.Id, u.FirstName, u.LastName, u.Email, u.Password)
}
