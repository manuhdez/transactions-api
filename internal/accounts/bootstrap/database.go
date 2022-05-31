package bootstrap

import (
	"database/sql"

	"github.com/manuhdez/transactions-api/internal/accounts/config"
)

func InitializeDB() *sql.DB {
	mysql, err := config.DBConnect()
	if err != nil {
		panic(err)
	}
	return mysql
}
