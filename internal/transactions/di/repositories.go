package di

import (
	"github.com/google/wire"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/transaction"
	"github.com/manuhdez/transactions-api/internal/transactions/infra"
)

var InitRepositories = wire.NewSet(
	wire.Bind(new(transaction.Repository), new(infra.TransactionMysqlRepository)),
	infra.NewTransactionMysqlRepository,
)
