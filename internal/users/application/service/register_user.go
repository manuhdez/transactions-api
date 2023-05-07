package service

import (
	"context"
	"log"

	"github.com/manuhdez/transactions-api/internal/users/domain/event"
	"github.com/manuhdez/transactions-api/internal/users/domain/service"
	"github.com/manuhdez/transactions-api/internal/users/domain/user"
)

type RegisterUser struct {
	repository user.Repository
	hasher     service.HashService
	eventBus   event.Bus
}

func NewRegisterUserService(repository user.Repository, hasher service.HashService, bus event.Bus) RegisterUser {
	return RegisterUser{
		repository: repository,
		hasher:     hasher,
		eventBus:   bus,
	}
}

func (srv RegisterUser) Register(user user.User) error {
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

	go srv.publishEvent(ctx, user)
	return nil
}

func (srv RegisterUser) publishEvent(ctx context.Context, user user.User) {
	ev := event.NewUserSignedUp(event.UserSignedUpBody{
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
