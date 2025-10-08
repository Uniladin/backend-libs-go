package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Publisher handles publishing messages to RabbitMQ
type Publisher struct {
	rabbitmq *RabbitMQ
}

// NewPublisher creates a new Publisher
func NewPublisher(rabbitmq *RabbitMQ) *Publisher {
	return &Publisher{
		rabbitmq: rabbitmq,
	}
}

// PublishMessage publishes a message to a specific exchange with a routing key
func (p *Publisher) PublishMessage(ctx context.Context, exchange, routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = p.rabbitmq.channel.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent, // Make message persistent
			Timestamp:    time.Now(),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	log.Printf("📤 Published message to exchange '%s' with routing key '%s'", exchange, routingKey)
	return nil
}

// PublishToQueue publishes a message directly to a queue (using default exchange)
func (p *Publisher) PublishToQueue(ctx context.Context, queueName string, message interface{}) error {
	return p.PublishMessage(ctx, "", queueName, message)
}

// PublishWithConfirm publishes a message with publisher confirmation
func (p *Publisher) PublishWithConfirm(ctx context.Context, exchange, routingKey string, message interface{}) error {
	// Enable publisher confirms
	if err := p.rabbitmq.channel.Confirm(false); err != nil {
		return fmt.Errorf("failed to enable publisher confirms: %w", err)
	}

	confirms := p.rabbitmq.channel.NotifyPublish(make(chan amqp.Confirmation, 1))

	body, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	err = p.rabbitmq.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         body,
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}

	// Wait for confirmation
	select {
	case confirm := <-confirms:
		if !confirm.Ack {
			return fmt.Errorf("message not confirmed by broker")
		}
		log.Printf("✅ Message confirmed by broker - exchange: '%s', routing key: '%s'", exchange, routingKey)
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(5 * time.Second):
		return fmt.Errorf("timeout waiting for confirmation")
	}

	return nil
}
