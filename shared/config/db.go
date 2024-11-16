package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Schema   string
}

func NewDBConfig() DBConfig {
	conf := DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
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

func NewDBConnection(c DBConfig) *sql.DB {
	log.Printf("[NewDBConnection][msg: connecting database %s]", c.Schema)

	dbUri := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.Schema,
	)

	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		panic(err)
	} else {
		log.Printf("[NewDBConnection][msg: connected database %s", c.Schema)
	}

	return db
}

func NewGormDBConnection(c DBConfig) *gorm.DB {
	log.Printf("[NewGormDBConnection][msg: connecting database %s]", c.Schema)

	dbUri := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		c.Host, c.Port, c.User, c.Password, c.Database, c.Schema,
	)

	db, err := gorm.Open(postgres.Open(dbUri))
	if err != nil {
		log.Printf("[NewGormDBConnection][msg: error connecting to database %s]", c.Schema)
		panic("cannot connect to database")
	} else {
		log.Printf("[NewGormDBConnection][msg: connected to database %s]", c.Schema)
	}

	return db
}
