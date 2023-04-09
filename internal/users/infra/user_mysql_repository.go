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

func (repo UserMysqlRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	row := repo.db.QueryRowContext(ctx, "select * from users where email = ?", email)

	var u UserMysql
	err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Password)
	if err != nil {
		log.Printf("unable to scan row into user variable: %e", err)
		return user.User{}, err
	}

	return user.New(u.Id, u.FirstName, u.LastName, u.Email, u.Password), nil
}
