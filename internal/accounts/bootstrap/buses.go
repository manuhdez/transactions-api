package bootstrap

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
	"github.com/manuhdez/transactions-api/internal/accounts/infra"
)

var InitBuses = wire.NewSet(
	wire.Bind(new(event.Bus), new(infra.EventBus)),
	infra.NewEventBus,
)
