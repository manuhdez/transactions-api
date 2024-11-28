package service

import (
	"context"
	"log"

	event2 "github.com/manuhdez/transactions-api/internal/users/internal/domain/event"
	"github.com/manuhdez/transactions-api/internal/users/internal/domain/service"
	user2 "github.com/manuhdez/transactions-api/internal/users/internal/domain/user"
)

type RegisterUser struct {
	repository user2.Repository
	hasher     domain_service.HashService
	eventBus   event2.Bus
}

func NewRegisterUserService(repository user2.Repository, hasher domain_service.HashService, bus event2.Bus) RegisterUser {
	return RegisterUser{
		repository: repository,
		hasher:     hasher,
		eventBus:   bus,
	}
}

func (srv RegisterUser) Register(user user2.User) error {
	ctx := context.Background()

	hashedPassword, err := srv.hasher.Hash(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	err = srv.repository.Save(ctx, user)
	if err != nil {
		return err
	}

	srv.publishEvent(ctx, user)
	return nil
}

func (srv RegisterUser) publishEvent(ctx context.Context, user user2.User) {
	ev := event2.NewUserSignedUp(event2.UserSignedUpBody{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	})

	err := srv.eventBus.Publish(ctx, ev)
	if err != nil {
		log.Println("error publishing user signed up event:", err)
	}
}
