package infra

import "github.com/manuhdez/transactions-api/internal/accounts/domain/event"

type Event struct {
	eventType event.Type
	body      []byte
}

func (m Event) Type() event.Type {
	return m.eventType
}

func (m Event) Body() []byte {
	return m.body
}
