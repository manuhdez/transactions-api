package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/transactions/config"
	"github.com/manuhdez/transactions-api/internal/transactions/domain/event"
	"github.com/streadway/amqp"
)

const (
	queueName     = "hello"
	queueConsumer = "transactions"
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

func (b EventBus) Listen() {
	channel, err := b.connection.Channel()
	if err != nil {
		log.Printf("Cannot connect to the rabbitmq channel: %e", err)
		return
	}

	messages, err := channel.Consume(queueName, queueConsumer, false, false, false, false, nil)
	if err != nil {
		log.Printf("Error consuming queued messages: %e", err)
		return
	}

	var forever chan struct{}

	go b.handleMessages(messages)

	log.Printf("Listening for enqueued messages")
	<-forever
}

type messageBody struct {
	Type string `json:"type"`
}

func (b EventBus) handleMessages(messages <-chan amqp.Delivery) {
	for message := range messages {

		var m messageBody
		err := json.Unmarshal(message.Body, &m)
		if err != nil {
			log.Printf("Error parsing message: %e", err)
			return
		}

		log.Printf("Received event with type: %s\n", m.Type)

		handler, ok := b.handlers[event.Type(m.Type)]
		if !ok {
			log.Printf("Handler for event type %s does not exist", m.Type)
			return
		}

		// Trigger the event handler
		err = handler.Handle(context.Background(), Event{event.Type(m.Type), message.Body})
		if err != nil {
			log.Printf("Error handling event `%s`: %e", m.Type, err)
			return
		}

		err = message.Ack(false)
		if err != nil {
			log.Printf("Error acknowledging event `%s`: %e", m.Type, err)
			return
		}
	}
}
