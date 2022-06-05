package config

import (
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

func NewDBConfig() DBConfig {
	conf := DBConfig{
		Host:     os.Getenv("MYSQL_HOST"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}

	// check if config has zero value
	if conf == (DBConfig{}) {
		panic("DB config is not set")
	}

	return conf
}
