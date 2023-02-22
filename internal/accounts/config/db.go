package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
}

func getDBConfig() DBConfig {
	conf := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}

	// check if config has zero value
	if conf == (DBConfig{}) {
		panic("DB config is not set")
	}

	return conf
}

func DBConnect() (*sql.DB, error) {
	c := getDBConfig()
	dbUri := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", c.User, c.Password, c.Host, c.Database)
	return sql.Open("mysql", dbUri)
}
