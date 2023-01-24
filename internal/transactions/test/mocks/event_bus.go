package mocks

import (
	"context"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/stretchr/testify/mock"
)

type EventBus struct {
	mock.Mock
}

func (m *EventBus) Publish(ctx context.Context, e event.Event) error {
	fmt.Println("event published:", e)
	args := m.Called(ctx, e)
	return args.Error(0)
}

func (m *EventBus) Subscribe(t event.Type, handler event.Handler) {
	m.Called(t, handler)
}
