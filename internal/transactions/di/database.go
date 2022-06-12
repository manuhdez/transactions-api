package di

import (
	"database/sql"
	"fmt"

	"github.com/manuhdez/transactions-api/internal/transactions/config"
)

func NewDBConnection() *sql.DB {
	c := config.NewDBConfig()
	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", c.User, c.Password, c.Host, c.Database)
	db, err := sql.Open("mysql", dbUri)
	if err != nil {
		panic(err)
	}
	return db
}
