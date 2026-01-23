package rabbitmq

import (
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Config represents RabbitMQ configuration
type Config struct {
	Addr string
	User string
	Pass string
}

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQ creates a new RabbitMQ connection
func NewRabbitMQ(cfg Config) (*RabbitMQ, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s/",
		cfg.User,
		cfg.Pass,
		cfg.Addr,
	)

	var conn *amqp.Connection
	var err error

	// Retry connection with exponential backoff
	maxRetries := DefaultMaxRetries
	for i := 0; i < maxRetries; i++ {
		conn, err = amqp.Dial(url)
		if err == nil {
			break
		}

		waitTime := time.Duration(i+1) * time.Second
		log.Printf("Failed to connect to RabbitMQ (attempt %d/%d): %v. Retrying in %v...",
			i+1, maxRetries, err, waitTime)
		time.Sleep(waitTime)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ after %d attempts: %w", maxRetries, err)
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open channel: %w", err)
	}

	log.Println("✅ Successfully connected to RabbitMQ")

	mq := RabbitMQ{
		conn:    conn,
		channel: channel,
	}

	err = mq.setupExchanges()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to setup exchanges: %w", err)
	}

	return &mq, nil
}

// Close closes the RabbitMQ connection and channel
func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
	log.Println("RabbitMQ connection closed")
}

// GetChannel returns the RabbitMQ channel
func (r *RabbitMQ) GetChannel() *amqp.Channel {
	return r.channel
}

// GetConnection returns the RabbitMQ connection
func (r *RabbitMQ) GetConnection() *amqp.Connection {
	return r.conn
}

// DeclareExchange declares an exchange
func (r *RabbitMQ) DeclareExchange(name, kind string, durable, autoDelete bool) error {
	return r.channel.ExchangeDeclare(
		name,       // name
		kind,       // type (direct, topic, fanout, headers)
		durable,    // durable
		autoDelete, // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
}

// DeclareQueue declares a queue
func (r *RabbitMQ) DeclareQueue(name string, durable, autoDelete bool) (amqp.Queue, error) {
	return r.channel.QueueDeclare(
		name,       // name
		durable,    // durable
		autoDelete, // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
}

// BindQueue binds a queue to an exchange with a routing key
func (r *RabbitMQ) BindQueue(queueName, routingKey, exchangeName string) error {
	return r.channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,        // no-wait
		nil,          // arguments
	)
}
