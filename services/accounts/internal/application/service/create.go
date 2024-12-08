package service

import (
	"context"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/transaction"
)

type CreateService struct {
	repository account.Repository
	bus        event.Bus
}

func NewCreateService(repository account.Repository, bus event.Bus) CreateService {
	return CreateService{repository, bus}
}

// Create creates a new account, store it in database and publish a domain event
func (s CreateService) Create(ctx context.Context, acc account.Account) error {
	if err := s.repository.Create(ctx, acc); err != nil {
		return fmt.Errorf("[CreateService:Create][err: %w]", err)
	}

	if err := s.bus.Publish(ctx, transaction.NewAccountCreated(acc)); err != nil {
		log.Printf("[CreateService:Create][event: NewAccountCreated][err: %s]", err)
	}

	return nil
}
