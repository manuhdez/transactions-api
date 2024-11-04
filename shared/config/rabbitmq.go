package config

import (
	"log"
	"os"
)

type RabbitMQConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

// NewRabbitMQConfig returns a config object with rabbitmq connection configuration
func NewRabbitMQConfig() RabbitMQConfig {
	conf := RabbitMQConfig{
		Host:     os.Getenv("AMQP_HOST"),
		Port:     os.Getenv("AMQP_PORT"),
		User:     os.Getenv("AMQP_USER"),
		Password: os.Getenv("AMQP_PASSWORD"),
	}
	if conf == (RabbitMQConfig{}) {
		log.Fatal("RabbitMQ config is not set")
	}
	return conf
}
