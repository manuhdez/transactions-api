package infra

import (
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/account"
)

type AccountMysqlRepository struct {
	db *gorm.DB
}

func NewAccountMysqlRepository(db *gorm.DB) AccountMysqlRepository {
	return AccountMysqlRepository{db: db}
}

// Save saves a new account
func (r AccountMysqlRepository) Save(ctx context.Context, acc account.Account) error {
	log.Printf("[AccountMysqlRepository:Save][account:%+v]", acc)

	res := r.db.WithContext(ctx).Create(&AccountMysql{
		Id:     acc.Id,
		UserId: acc.UserId,
	})
	if res.Error != nil {
		return fmt.Errorf("[AccountMysqlRepository:Save][err: %w]", res.Error)
	}

	return nil
}

// FindById returns an account found by id
func (r AccountMysqlRepository) FindById(ctx context.Context, id string) (account.Account, error) {
	log.Printf("[AccountMysqlRepository:FindById][id:%s]", id)

	var acc AccountMysql
	res := r.db.WithContext(ctx).Model(&AccountMysql{}).Where("id = ?", id).First(&acc)
	if res.Error != nil {
		return account.Account{}, fmt.Errorf("[AccountMysqlRepository:FindById][id:%s][err:%w]", id, res.Error)
	}

	return acc.ToDomainModel(), nil
}
