package infra

import "github.com/manuhdez/transactions-api/internal/users/domain/event"

type Event struct {
	eventType event.Type
	body      []byte
}

func (e Event) Type() event.Type {
	return e.eventType
}

func (e Event) Body() []byte {
	return e.body
}
