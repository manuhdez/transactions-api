package service

import (
	"context"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
)

type TransactionsRetriever struct {
	repository transaction.Repository
}

func NewTransactionsRetriever(r transaction.Repository) TransactionsRetriever {
	return TransactionsRetriever{r}
}

// Retrieve returns a list with all the transactions associated to a given user
func (s TransactionsRetriever) Retrieve(ctx context.Context, userId string) ([]transaction.Transaction, error) {
	log.Printf("[TransactionsRetriever:Retrieve][userId:%s]", userId)

	trxList, err := s.repository.All(ctx, userId)
	if err != nil {
		log.Printf("[TransactionsRetriever:Retrieve][userId:%s][err:%s]", userId, err)
		return nil, fmt.Errorf("[TransactionsRetriever:Retrieve][err: database error]")
	}

	return trxList, nil
}
