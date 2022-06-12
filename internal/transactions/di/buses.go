package di

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/manuhdez/transactions-api/internal/transactions/infra"
)

var InitBuses = wire.NewSet(
	wire.Bind(new(event.Bus), new(infra.EventBus)),
	infra.NewEventBus,
)
