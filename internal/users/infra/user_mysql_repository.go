package infra

import (
	"context"
	"database/sql"
	"log"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type UserMysqlRepository struct {
	db *sql.DB
}

func NewUserMysqlRepository(db *sql.DB) UserMysqlRepository {
	return UserMysqlRepository{db: db}
}

func (repo UserMysqlRepository) Save(ctx context.Context, u user.User) error {
	_, err := repo.db.ExecContext(
		ctx,
		"insert into users (id, first_name, last_name, email, password) values (?, ?, ?, ?, ?);",
		u.Id, u.FirstName, u.LastName, u.Email, u.Password,
	)

	if err != nil {
		log.Printf("Error saving new user: %e", err)
	}
	return err
}
