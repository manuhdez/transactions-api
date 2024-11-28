package service

import (
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
)

type EventPuller interface {
	PullEvents() []event.Event
}
