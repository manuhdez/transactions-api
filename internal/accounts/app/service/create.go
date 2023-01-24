package service

import (
	"context"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/domain/account"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
)

type CreateService struct {
	repository account.Repository
	bus        event.Bus
}

func NewCreateService(repository account.Repository, bus event.Bus) CreateService {
	return CreateService{repository, bus}
}

func (s CreateService) Create(a account.Account) error {
	err := s.repository.Create(a)
	if err != nil {
		return err
	}

	go func() {
		err := s.bus.Publish(context.Background(), event.NewAccountCreated(a.Id(), a.Balance()))
		if err != nil {
			log.Println("error publishing account created event:", err)
		}
	}()

	return nil
}
