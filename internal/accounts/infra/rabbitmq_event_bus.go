package infra

import (
	"context"
	"fmt"
	"log"

	"github.com/manuhdez/transactions-api/internal/accounts/config"
	"github.com/manuhdez/transactions-api/internal/accounts/domain/event"
	"github.com/streadway/amqp"
)

const (
	exchangeName  = "transactions-api-exchange"
	queueName     = "accounts-queue"
	queueConsumer = "accounts"
)

var routingKeys = []string{
	"event.transactions.*",
	"event.users.*",
}

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

	err = ch.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to create exchange: %e", err)
	}

	_, err = ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		return err
	}

	err = bindQueues(ch)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func bindQueues(ch *amqp.Channel) error {
	var errCount int
	var e string
	for _, routingKey := range routingKeys {
		err := ch.QueueBind(queueName, routingKey, exchangeName, false, nil)
		if err != nil {
			e += fmt.Sprintf("Failed to bind queue `%s`: %e\n", queueName, err)
		} else {
			log.Printf("Bound queue `%s` to routing key `%s`", queueName, routingKey)
		}
	}

	if errCount > 0 {
		return fmt.Errorf("%s", e)
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

	key := string(event.Type())
	err = ch.Publish(exchangeName, key, false, false, amqp.Publishing{
		Type:        key,
		ContentType: "text/plain",
		Body:        event.Body(),
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
	for msg := range messages {
		log.Printf("Received event with type: %s\n", msg.Type)

		handler, ok := b.handlers[event.Type(msg.Type)]
		if !ok {
			log.Printf("Handler for event type %s does not exist", msg.Type)
			return
		}

		err := handler.Handle(context.Background(), Event{event.Type(msg.Type), msg.Body})
		if err != nil {
			log.Printf("error handling event %s: %e", msg.Type, err)
			return
		}

		err = msg.Ack(false)
		if err != nil {
			log.Printf("error acknowledging message: %e", err)
			return
		}
	}
}
