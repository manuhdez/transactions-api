package rabbitmq

import (
	"context"
	"fmt"
	"log"

	"github.com/streadway/amqp"

	"github.com/manuhdez/transactions-api/internal/users/internal/domain/event"
	"github.com/manuhdez/transactions-api/shared/config"
)

const (
	exchangeName  = "transactions-api-exchange"
	routingKey    = "event.users.*"
	queueName     = "users-queue"
	queueConsumer = "users"
)

type RabbitEventBus struct {
	connection *amqp.Connection
	handlers   map[event.Type]event.Handler
}

func NewAmqpConnection() (*amqp.Connection, error) {
	conf := config.NewRabbitMQConfig()
	uri := fmt.Sprintf("amqp://%s:%s@%s:%s/", conf.User, conf.Password, conf.Host, conf.Port)
	return amqp.Dial(uri)
}

func NewRabbitEventBus() RabbitEventBus {
	con, err := NewAmqpConnection()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}

	channel, err := con.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel to RabbitMQ: %e", err)
	}

	err = channel.ExchangeDeclare(exchangeName, "topic", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to create exchange: %e", err)
	}

	_, err = channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare queue %s: %e", queueName, err)
	}

	err = channel.QueueBind(queueName, routingKey, exchangeName, false, nil)
	if err != nil {
		log.Fatalf("Failed to bind queue `%s` to the exchange `%s`: %e", queueName, exchangeName, err)
	}

	return RabbitEventBus{
		connection: con,
		handlers:   RegisterHandlers(),
	}
}

func RegisterHandlers(eventHandlers ...event.Handler) map[event.Type]event.Handler {
	h := make(map[event.Type]event.Handler)
	for _, eventHandler := range eventHandlers {
		h[eventHandler.Type()] = eventHandler
	}
	return h
}

func (b RabbitEventBus) Publish(_ context.Context, event event.Event) error {
	ch, err := b.connection.Channel()
	if err != nil {
		return fmt.Errorf("[RabbitEventBus:Publish][msg: cannot stablish channel connection][err:%w]", err)
	}

	defer ch.Close()

	key := event.Type().String()
	if err = ch.Publish(exchangeName, key, false, false, amqp.Publishing{
		Type:        key,
		Body:        event.Body(),
		ContentType: "text/plain",
	}); err != nil {
		return fmt.Errorf("[RabbitEventBus:Publish][msg: could not publish message %s][err: %w]", key, err)
	}

	log.Printf("Message %s sent!", key)
	return nil
}

func (b RabbitEventBus) Subscribe(t event.Type, h event.Handler) {
	b.handlers[t] = h
}

func (b RabbitEventBus) Listen() {
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

func (b RabbitEventBus) handleMessages(messages <-chan amqp.Delivery) {
	for msg := range messages {
		log.Printf("Received event with type: %s\n", msg.Type)

		handler, ok := b.handlers[event.Type(msg.Type)]
		if !ok {
			errMsg := fmt.Errorf("%w: %s", event.ErrEventNotFound, msg.Type)
			log.Printf(errMsg.Error())
			return
		}

		err := handler.Handle(context.Background(), Event{event.Type(msg.Type), msg.Body})
		if err != nil {
			log.Printf("❌ There was an error while handling the event `%s`: %e", msg.Type, err)
			return
		}

		err = msg.Ack(false)
		if err != nil {
			log.Printf("❌ Cannot ack event `%s`: %e", msg.Type, err)
			return
		}
	}
}
