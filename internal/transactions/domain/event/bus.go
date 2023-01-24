package event

import (
	"context"
	"errors"
)

type Bus interface {
	Publish(ctx context.Context, event Event) error
	Subscribe(Type, Handler)
	Listen()
}

//go:generate mockery --case=snake --outpkg=mocks --output=../../test/mocks --name=Bus

type Type string

type Event interface {
	Type() Type
	Body() []byte
}

var (
	ErrEventNotFound = errors.New("event not found")
)

type Handler interface {
	Handle(ctx context.Context, event Event) error
}
