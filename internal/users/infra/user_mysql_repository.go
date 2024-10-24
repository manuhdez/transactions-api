package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type UserMysqlRepository struct {
	db *sql.DB
}

func NewUserMysqlRepository(db *sql.DB) UserMysqlRepository {
	return UserMysqlRepository{db: db}
}

func (repo UserMysqlRepository) All(ctx context.Context) ([]user.User, error) {
	rows, err := repo.db.QueryContext(ctx, "select * from users")
	if err != nil {
		log.Printf("Error querying users: %e", err)
		return nil, err
	}

	var users []user.User
	for rows.Next() {
		var u UserMysql
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Password)
		if err != nil {
			log.Printf("Error scanning row: %e", err)
			return nil, err
		}
		users = append(users, u.ToDomainModel())
	}

	return users, nil
}

func (repo UserMysqlRepository) Save(ctx context.Context, u user.User) error {
	if err := repo.ping(); err != nil {
		log.Printf("[UserMysqlRepository:Save]%s", err)
		return fmt.Errorf("[UserMysqlRepository:Save][err: cannot connect to database]")
	}

	_, err := repo.db.ExecContext(
		ctx,
		"insert into users (id, first_name, last_name, email, password) values ($1, $2, $3, $4, $5);",
		u.Id, u.FirstName, u.LastName, u.Email, u.Password,
	)

	if err != nil {
		er := fmt.Errorf("[UserMysqlRepository:Save][err: %w]", err)
		log.Println(er)
		return err
	}

	return nil
}

func (repo UserMysqlRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	row := repo.db.QueryRowContext(ctx, "select * from users where email = $1", email)

	var u UserMysql
	err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Password)
	if err != nil {
		log.Printf("unable to scan row into user variable: %e", err)
		return user.User{}, err
	}

	return user.New(u.Id, u.FirstName, u.LastName, u.Email, u.Password), nil
}

func (repo UserMysqlRepository) ping() error {
	if err := repo.db.Ping(); err != nil {
		return fmt.Errorf("[ping][err: %w]", err)
	}

	return nil
}
