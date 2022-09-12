package infra

import (
	"context"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/transactions/config"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/streadway/amqp"
)

type EventBus struct {
	connection *amqp.Connection
	handlers   map[event.Type]event.Handler
}

func NewAmqpConnection() (*amqp.Connection, error) {
	conf := config.NewRabbitMQConfig()
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port)
	return amqp.Dial(uri)
}

func NewEventBus() EventBus {
	con, err := NewAmqpConnection()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	return EventBus{con, make(map[event.Type]event.Handler)}
}

func (b EventBus) Publish(_ context.Context, event event.Event) error {
	ch, err := b.connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Event with type: %s and body: %s", event.Type(), event.Body())
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		Type:        string(event.Type()),
		ContentType: "text/plain",
		Body:        event.Body(),
	})

	if err != nil {
		return err
	}
	log.Printf("Message %s sent!", event.Type())
	return nil
}

func (b EventBus) Subscribe(t event.Type, h event.Handler) {
	b.handlers[t] = h
}
