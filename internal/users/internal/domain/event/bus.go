package event

import (
	"context"
	"errors"
)

var (
	ErrEventNotFound = errors.New("event not found")
)

type Type string

func (t Type) String() string {
	return string(t)
}

type Event interface {
	Type() Type
	Body() []byte
}

type Bus interface {
	Publish(ctx context.Context, event Event) error
	Subscribe(Type, Handler)
	Listen()
}

type Handler interface {
	Type() Type
	Handle(ctx context.Context, event Event) error
}
