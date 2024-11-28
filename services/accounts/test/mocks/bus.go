package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"

	"github.com/manuhdez/transactions-api/internal/accounts/internal/domain/event"
)

type EventBus struct {
	mock.Mock
}

func (m *EventBus) Publish(ctx context.Context, e event.Event) error {
	args := m.Called(ctx, e)
	return args.Error(0)
}

func (m *EventBus) Subscribe(t event.Type, handler event.Handler) {
	m.Called(t, handler)
}

func (m *EventBus) Listen() {}
