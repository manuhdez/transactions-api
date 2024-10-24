package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Schema   string
}

func newDBConfig() DBConfig {
	conf := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
		Schema:   os.Getenv("DB_SCHEMA"),
	}

	// check if config has zero value
	if conf == (DBConfig{}) {
		panic("DB config is not set")
	}

	return conf
}

func NewDBConnection() *sql.DB {
	c := newDBConfig()
	dbUri := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s", c.Host, c.Port, c.User, c.Password, c.Database, c.Schema)
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	}
	return db
}
