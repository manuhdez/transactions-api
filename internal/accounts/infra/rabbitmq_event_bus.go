package infra

import (
	"context"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/app/handler"
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

func readMessages(con *amqp.Connection) {
	c, err := con.Channel()
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

	go func() {
		for d := range messages {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}

func getEventHandlers() map[event.Type]event.Handler {
	var handlers = make(map[event.Type]event.Handler)
	handlers[handler.DepositCreatedType] = handler.DepositCreated{}
	return handlers
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

	go readMessages(con)

	return EventBus{con, getEventHandlers()}
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
