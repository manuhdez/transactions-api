package bootstrap

import (
	"database/sql"

	"github.com/manuhdez/transactions-api/internal/accounts/config"
)

type Databases struct {
	Mysql *sql.DB
}

func bootstrapDatabases() Databases {
	mysql, err := config.DBConnect()
	if err != nil {
		panic(err)
	}

	return Databases{
		Mysql: mysql,
	}
}
