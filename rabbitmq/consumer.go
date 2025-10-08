package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// MessageHandler is a function type for handling consumed messages
type MessageHandler func(ctx context.Context, body []byte) error

// Consumer handles consuming messages from RabbitMQ
type Consumer struct {
	rabbitmq *RabbitMQ
}

// NewConsumer creates a new Consumer
func NewConsumer(rabbitmq *RabbitMQ) *Consumer {
	return &Consumer{
		rabbitmq: rabbitmq,
	}
}

// ConsumeQueue consumes messages from a queue
func (c *Consumer) ConsumeQueue(ctx context.Context, queueName, consumerTag string, handler MessageHandler) error {
	msgs, err := c.rabbitmq.channel.Consume(
		queueName,   // queue
		consumerTag, // consumer tag
		false,       // auto-ack (set to false for manual acknowledgment)
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return fmt.Errorf("failed to register consumer: %w", err)
	}

	log.Printf("📥 Started consuming from queue '%s' with tag '%s'", queueName, consumerTag)

	go func() {
		for {
			select {
			case <-ctx.Done():
				log.Printf("Stopping consumer '%s' for queue '%s'", consumerTag, queueName)
				return
			case msg, ok := <-msgs:
				if !ok {
					log.Printf("Consumer channel closed for queue '%s'", queueName)
					return
				}

				// Process the message
				if err := handler(ctx, msg.Body); err != nil {
					log.Printf("❌ Error handling message from queue '%s': %v", queueName, err)
					if IsUnrecoverableError(err) {
						// Positive acknowledgment
						msg.Ack(false)
					} else {
						// Negative acknowledgment - requeue the message
						msg.Nack(false, true)
					}
				} else {
					// Positive acknowledgment
					msg.Ack(false)
					log.Printf("✅ Successfully processed message from queue '%s'", queueName)
				}
			}
		}
	}()

	return nil
}

// ConsumeWithPrefetch consumes messages with a prefetch count (QoS)
func (c *Consumer) ConsumeWithPrefetch(ctx context.Context, queueName, consumerTag string, prefetchCount int, handler MessageHandler) error {
	// Set QoS (Quality of Service)
	err := c.rabbitmq.channel.Qos(
		prefetchCount, // prefetch count
		0,             // prefetch size
		false,         // global
	)
	if err != nil {
		return fmt.Errorf("failed to set QoS: %w", err)
	}

	return c.ConsumeQueue(ctx, queueName, consumerTag, handler)
}

// UnmarshalMessage is a helper function to unmarshal JSON message body
func UnmarshalMessage(body []byte, v interface{}) error {
	if err := json.Unmarshal(body, v); err != nil {
		return fmt.Errorf("failed to unmarshal message: %w", err)
	}
	return nil
}

func IsUnrecoverableError(err error) bool {
	errMsg := err.Error()
	// List of unrecoverable error indicators
	unrecoverableIndicators := []string{
		"json: cannot unmarshal", // JSON unmarshal errors
		"failed to unmarshal",    // Custom unmarshal errors
		"invalid argument",       // Invalid argument errors
		// Add more indicators as needed
	}
	// Implement logic to determine if the error is unrecoverable
	// For example, check for specific error types or messages
	for _, indicator := range unrecoverableIndicators {
		if strings.Contains(errMsg, indicator) {
			return true
		}
	}
	return false
}
