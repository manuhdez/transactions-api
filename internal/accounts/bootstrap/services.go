package bootstrap

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/accounts/app/service"
)

var InitServices = wire.NewSet(
	service.NewCreateService,
	service.NewFindAllService,
	service.NewFindAccountService,
	service.NewDeleteAccountService,
)
