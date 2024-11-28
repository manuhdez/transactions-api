package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
)

type TransactionMysqlRepository struct {
	db *gorm.DB
}

func NewTransactionMysqlRepository(db *gorm.DB) TransactionMysqlRepository {
	return TransactionMysqlRepository{
		db: db,
	}
}

// Deposit saves a deposit transaction
func (r TransactionMysqlRepository) Deposit(ctx context.Context, trx transaction.Transaction) error {
	log.Printf("[TransactionMysqlRepository:Deposit][trx:%+v]", trx)
	if err := r.saveTransaction(ctx, trx); err != nil {
		return fmt.Errorf("[TransactionMysqlRepository:Deposit]%w", err)
	}
	return nil
}

// Withdraw saves a withdraw transaction
func (r TransactionMysqlRepository) Withdraw(ctx context.Context, trx transaction.Transaction) error {
	log.Printf("[TransactionMysqlRepository:Withdraw][trx:%+v]", trx)
	if err := r.saveTransaction(ctx, trx); err != nil {
		return fmt.Errorf("[TransactionMysqlRepository:Withdraw]%w", err)
	}
	return nil
}

// All retrieves a list with all transactions
func (r TransactionMysqlRepository) All(ctx context.Context, userId string) ([]transaction.Transaction, error) {
	log.Printf("[TransactionMysqlRepository:All][userId:%s]", userId)

	var trxList []TransactionMysql
	res := r.db.
		WithContext(ctx).
		Model(&TransactionMysql{}).
		Where("user_id = ?", userId).
		Order("date desc").
		Find(&trxList)
	if res.Error != nil {
		log.Printf("[TransactionMysqlRepository:All][gorm][err:%s]", res.Error)
		return nil, fmt.Errorf("[TransactionMysqlRepository:All][gorm][err:%w]", res.Error)
	}

	transactions := make([]transaction.Transaction, len(trxList))
	for i := range trxList {
		transactions[i] = trxList[i].ToDomainModel()
	}

	return transactions, nil
}

// FindByAccount retrieves the transaction list from a given account
func (r TransactionMysqlRepository) FindByAccount(ctx context.Context, id string) ([]transaction.Transaction, error) {
	log.Printf("[TransactionMysqlRepository:FindByAccount][accountId:%s]", id)

	var trxSQL []TransactionMysql
	res := r.db.
		WithContext(ctx).
		Model(&TransactionMysql{}).
		Where("account_id = ?", id).
		Find(&trxSQL)

	if res.Error != nil {
		log.Printf("[TransactionMysqlRepository:FindByAccount][accountId:%s][err:%s]", id, res.Error)
		return nil, fmt.Errorf("[TransactionMysqlRepository:FindByAccount][accountId:%s][err: database error]", id)
	}

	transactions := make([]transaction.Transaction, len(trxSQL))
	for i := range trxSQL {
		transactions[i] = trxSQL[i].ToDomainModel()
	}

	return transactions, nil
}

// saveTransaction saves a transaction in the database
func (r TransactionMysqlRepository) saveTransaction(ctx context.Context, trx transaction.Transaction) error {
	res := r.db.WithContext(ctx).Create(&TransactionMysql{
		AccountId: trx.AccountId,
		UserId:    trx.UserId,
		Amount:    trx.Amount,
		Balance:   trx.Amount, // TODO: calculate balance correctly
		Type:      trx.Type,
		Date:      time.Now(),
	})
	if res.Error != nil {
		return fmt.Errorf("[saveTransaction][err: %w]", res.Error)
	}

	return nil
}
