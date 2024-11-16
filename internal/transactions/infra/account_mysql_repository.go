package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
)

type AccountMysqlRepository struct {
	db *sql.DB
}

func NewAccountMysqlRepository(db *sql.DB) AccountMysqlRepository {
	return AccountMysqlRepository{db: db}
}

func (repo AccountMysqlRepository) Save(ctx context.Context, account account.Account) error {
	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO accounts (id, user_id) VALUES ($1, $2)",
		account.Id, account.UserId,
	)

	if err != nil {
		log.Printf("Error while saving new account")
		return err
	}
	return nil
}

// FindById returns an account found by id
func (repo AccountMysqlRepository) FindById(ctx context.Context, id string) (account.Account, error) {
	log.Printf("[AccountMysqlRepository:FindById][id:%s]", id)

	var acc AccountMysql
	err := repo.db.
		QueryRowContext(ctx, "select id, user_id from accounts where id=$1", id).
		Scan(&acc.Id, &acc.UserId)

	if err != nil {
		log.Printf("[AccountMysqlRepository:FindById][id:%s][err:%s]", id, err)
		return account.Account{}, fmt.Errorf("[AccountMysqlRepository:FindById][id:%s][err: database error]", id)
	}

	return acc.ToDomainModel(), nil
}
