package db

import (
	"context"
	"fmt"
	"log"

	"gorm.io/gorm"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
)

type AccountPostgresRepository struct {
	db *gorm.DB
}

func NewAccountPostgresRepository(db *gorm.DB) AccountPostgresRepository {
	return AccountPostgresRepository{db: db}
}

// Create saves a new account into db
func (r AccountPostgresRepository) Create(ctx context.Context, a account.Account) error {
	log.Printf("[AccountPostgresRepository:Create][account:%+v]", a)

	if err := r.db.WithContext(ctx).Create(&AccountPostgres{
		Id:       a.Id(),
		UserId:   a.UserId.String(),
		Balance:  a.Balance(),
		Currency: a.Currency(),
	}).Error; err != nil {
		// metrics.TrackDBErrorAdd()
		return fmt.Errorf("[AccountPostgres:Create][err: %w]", err)
	}

	return nil
}

// Find finds an account by id
func (r AccountPostgresRepository) Find(ctx context.Context, id string) (account.Account, error) {
	log.Printf("[AccountPostgresRepository:Find][accountId:%s]", id)

	var acc AccountPostgres
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&acc).Error; err != nil {
		// metrics.TrackDBErrorAdd()
		return account.Account{}, fmt.Errorf("[AccountPostgresRepository:Find][err: %w]", err)
	}

	return acc.parseToDomainModel(), nil
}

// GetByUserId returns the list of accounts for a given user
func (r AccountPostgresRepository) GetByUserId(ctx context.Context, userId string) ([]account.Account, error) {
	log.Printf("[AccountPostgresRepository:GetByUserId][userId:%s]", userId)

	var accounts []AccountPostgres
	if err := r.db.WithContext(ctx).Where("user_id = ?", userId).Find(&accounts).Error; err != nil {
		// metrics.TrackDBErrorAdd()
		return nil, fmt.Errorf("[AccountPostgresRepository:GetByUserId][err: %w]", err)
	}

	return parseToDomainModels(accounts), nil
}

// Delete deletes an account by id
func (r AccountPostgresRepository) Delete(ctx context.Context, id string) error {
	log.Printf("[AccountPostgresRepository:Delete][accountId:%s]", id)

	if err := r.db.WithContext(ctx).Delete(&AccountPostgres{Id: id}).Error; err != nil {
		// metrics.TrackDBErrorAdd()
		return fmt.Errorf("[AccountPostgresRepository:Delete][err: %w]", err)
	}

	return nil
}

// UpdateBalance updates the balance of an account
func (r AccountPostgresRepository) UpdateBalance(ctx context.Context, id string, balance float32) error {
	log.Printf("[AccountPostgresRepository:UpdateBalance][accountId:%s][balance:%f]", id, balance)

	if err := r.db.WithContext(ctx).Model(&AccountPostgres{}).Where("id = ?", id).Update("balance", balance).Error; err != nil {
		// metrics.TrackDBErrorAdd()
		return fmt.Errorf("[AccountPostgresRepository:UpdateBalance][err: %w]", err)
	}

	return nil
}
