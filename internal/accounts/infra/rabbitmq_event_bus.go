package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/config"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
	"github.com/streadway/amqp"
)

const (
	queueName     = "hello"
	queueConsumer = "accounts"
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

func createDefaultQueue(con *amqp.Connection) error {
	ch, err := con.Channel()
	if err != nil {
		return err
	}

	_, err = ch.QueueDeclare("hello", false, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func NewEventBus() EventBus {
	con, err := NewAmqpConnection()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %e", err)
	}

	err = createDefaultQueue(con)
	if err != nil {
		log.Fatalf("Failed to create message queue: %e", err)
	}

	return EventBus{con, make(map[event.Type]event.Handler)}
}

func (b EventBus) Publish(_ context.Context, event event.Event) error {
	ch, err := b.connection.Channel()
	if err != nil {
		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	body := event.Type()

	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(body),
	})
	if err != nil {
		return err
	}

	log.Printf("Message %s sent", event.Type())
	return nil
}

func (b EventBus) Subscribe(t event.Type, h event.Handler) {
	b.handlers[t] = h
}

func (b EventBus) Listen() {
	c, err := b.connection.Channel()
	if err != nil {
		log.Printf("Cannot consume queued messages: %e", err)
		return
	}

	messages, err := c.Consume(queueName, queueConsumer, false, false, false, false, nil)
	if err != nil {
		log.Printf("Cannot consume queued messages: %e", err)
		return
	}

	var forever chan struct{}

	go b.handleMessages(messages)

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}

type messageBody struct {
	Type string `json:"type"`
}

func (b EventBus) handleMessages(messages <-chan amqp.Delivery) {

	for d := range messages {

		var m messageBody
		e := json.Unmarshal(d.Body, &m)
		if e != nil {
			log.Printf("Error parsing message: %e", e)
		}

		log.Printf("Received a message from with type: %s", m.Type)

		h, ok := b.handlers[event.Type(m.Type)]
		if ok != true {
			log.Printf("handler not ok")
			return

		}

		_ = h.Handle(context.Background(), Event{event.Type(m.Type), d.Body})
	}
}
