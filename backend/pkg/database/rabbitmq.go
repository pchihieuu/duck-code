package database

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"backend/internal/config"
)

// NewRabbitMQ opens a connection + channel to RabbitMQ, used by the
// execution module to publish/consume Code Runner submission jobs.
// Kept separate from Redis (which stays cache + session per the stack).
func NewRabbitMQ(cfg *config.Config) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, nil, err
	}

	// Declare the queue up front so publishers/consumers never race on
	// "queue does not exist" during early development.
	_, err = ch.QueueDeclare(
		"code_execution_jobs", // name
		true,                  // durable
		false,                 // auto-delete
		false,                 // exclusive
		false,                 // no-wait
		nil,                   // args
	)
	if err != nil {
		ch.Close()
		conn.Close()
		return nil, nil, err
	}

	return conn, ch, nil
}
