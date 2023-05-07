package event

import (
	"context"
	"errors"
)

var (
	ErrEventNotFound = errors.New("event not found")
)

type Type string

type Event interface {
	Type() Type
	Body() []byte
}

//go:generate mockery --case=snake --outpkg=mocks --output=test/mocks --name=Bus

type Bus interface {
	Publish(ctx context.Context, event Event) error
	Subscribe(Type, Handler)
	Listen()
}

type Handler interface {
	Type() Type
	Handle(ctx context.Context, event Event) error
}
