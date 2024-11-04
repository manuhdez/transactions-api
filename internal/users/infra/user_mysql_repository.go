package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/manuhdez/transactions-api/internal/users/domain/user"
	"github.com/manuhdez/transactions-api/internal/users/infra/metrics"
)

type UserMysqlRepository struct {
	db *sql.DB
}

func NewUserMysqlRepository(db *sql.DB) UserMysqlRepository {
	return UserMysqlRepository{db: db}
}

func (repo UserMysqlRepository) All(ctx context.Context) ([]user.User, error) {
	defer metrics.TrackDBQueryDuration(time.Now())

	rows, err := repo.db.QueryContext(ctx, "select * from users")
	if err != nil {
		log.Printf("Error querying users: %e", err)
		metrics.TrackDBErrorAdd()
		return nil, err
	}

	var users []user.User
	for rows.Next() {
		var u UserMysql
		err := rows.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Password)
		if err != nil {
			log.Printf("Error scanning row: %e", err)
			metrics.TrackDBErrorAdd()
			return nil, err
		}
		users = append(users, u.ToDomainModel())
	}

	return users, nil
}

func (repo UserMysqlRepository) Save(ctx context.Context, u user.User) error {
	defer metrics.TrackDBQueryDuration(time.Now())

	if _, err := repo.db.ExecContext(
		ctx,
		"insert into users (id, first_name, last_name, email, password) values ($1, $2, $3, $4, $5);",
		u.Id, u.FirstName, u.LastName, u.Email, u.Password,
	); err != nil {
		er := fmt.Errorf("[UserMysqlRepository:Save][err: %w]", err)
		log.Println(er)
		metrics.TrackDBErrorAdd()
		return err
	}

	return nil
}

func (repo UserMysqlRepository) FindByEmail(ctx context.Context, email string) (user.User, error) {
	defer metrics.TrackDBQueryDuration(time.Now())

	row := repo.db.QueryRowContext(ctx, "select * from users where email = $1", email)
	if row.Err() != nil {
		er := fmt.Errorf("[UserMysqlRepository:FindByEmail][err: %w]", row.Err())
		log.Println(er)
		metrics.TrackDBErrorAdd()
		return user.User{}, er
	}

	var u UserMysql
	err := row.Scan(&u.Id, &u.FirstName, &u.LastName, &u.Email, &u.Password)
	if err != nil {
		er := fmt.Errorf("[UserMysqlRepository:FindByEmail][err: %w]", err)
		log.Println(er)
		metrics.TrackDBErrorAdd()
		return user.User{}, er
	}

	return user.New(u.Id, u.FirstName, u.LastName, u.Email, u.Password), nil
}
