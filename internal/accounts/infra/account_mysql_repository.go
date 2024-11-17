package infra

import (
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/users/infra/metrics"
)

type AccountMysqlRepository struct {
	db *gorm.DB
}

func NewAccountMysqlRepository(db *gorm.DB) AccountMysqlRepository {
	return AccountMysqlRepository{db: db}
}

// Create saves a new account into db
func (r AccountMysqlRepository) Create(ctx context.Context, a account.Account) error {
	log.Printf("[AccountMysqlRepository:Create][account:%+v]", a)

	if err := r.db.WithContext(ctx).Create(&AccountMysql{
		Id:       a.Id(),
		UserId:   a.UserId.String(),
		Balance:  a.Balance(),
		Currency: a.Currency(),
	}).Error; err != nil {
		metrics.TrackDBErrorAdd()
		return fmt.Errorf("[AccountMysql:Create][err: %w]", err)
	}

	return nil
}

// Find finds an account by id
func (r AccountMysqlRepository) Find(ctx context.Context, id string) (account.Account, error) {
	log.Printf("[AccountMysqlRepository:Find][accountId:%s]", id)

	var acc AccountMysql
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&acc).Error; err != nil {
		metrics.TrackDBErrorAdd()
		return account.Account{}, fmt.Errorf("[AccountMysqlRepository:Find][err: %w]", err)
	}

	return acc.parseToDomainModel(), nil
}

// GetByUserId returns the list of accounts for a given user
func (r AccountMysqlRepository) GetByUserId(ctx context.Context, userId string) ([]account.Account, error) {
	log.Printf("[AccountMysqlRepository:GetByUserId][userId:%s]", userId)

	var accounts []AccountMysql
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&accounts).Error; err != nil {
		metrics.TrackDBErrorAdd()
		return nil, fmt.Errorf("[AccountMysqlRepository:GetByUserId][err: %w]", err)
	}

	return parseToDomainModels(accounts), nil
}

// Delete deletes an account by id
func (r AccountMysqlRepository) Delete(ctx context.Context, id string) error {
	log.Printf("[AccountMysqlRepository:Delete][accountId:%s]", id)

	if err := r.db.WithContext(ctx).Delete(&AccountMysql{Id: id}).Error; err != nil {
		metrics.TrackDBErrorAdd()
		return fmt.Errorf("[AccountMysqlRepository:Delete][err: %w]", err)
	}

	return nil
}

// UpdateBalance updates the balance of an account
func (r AccountMysqlRepository) UpdateBalance(ctx context.Context, id string, balance float32) error {
	log.Printf("[AccountMysqlRepository:UpdateBalance][accountId:%s][balance:%f]", id, balance)

	if err := r.db.WithContext(ctx).Model(&AccountMysql{}).Where("id = ?", id).Update("balance", balance).Error; err != nil {
		metrics.TrackDBErrorAdd()
		return fmt.Errorf("[AccountMysqlRepository:UpdateBalance][err: %w]", err)
	}

	return nil
}
