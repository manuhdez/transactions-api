package infra

import (
	"context"
	"database/sql"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
	"log"
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
		"INSERT INTO accounts (id) VALUES (?)",
		account.Id,
	)

	if err != nil {
		log.Printf("Error while saving new account")
		return err
	}
	return nil
}
