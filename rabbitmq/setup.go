package rabbitmq

import (
	"log"
	"uniladin-backend-libs/rabbitmq/services/auth"
	"uniladin-backend-libs/rabbitmq/services/notification"
	"uniladin-backend-libs/rabbitmq/services/post"
	"uniladin-backend-libs/rabbitmq/services/profile"
)

// QueueConfig represents a queue configuration
type QueueConfig struct {
	Name       string
	RoutingKey string
	Exchange   string
}

type ExchangeConfig struct {
	Name string
	Type string // e.g., "topic", "direct", "fanout"
}

var exchangeConfigs []ExchangeConfig = []ExchangeConfig{
	{Name: auth.ExchangeAuthEvents, Type: ExchangeTypeDirect},
	{Name: profile.ExchangeProfileEvents, Type: ExchangeTypeDirect},
	{Name: notification.ExchangeNotificationEvents, Type: ExchangeTypeDirect},
	{Name: post.ExchangePostEvents, Type: ExchangeTypeDirect},
}

func (rmq *RabbitMQ) setupExchanges() error {
	for _, ec := range exchangeConfigs {
		err := rmq.DeclareExchange(ec.Name, ec.Type, true, false)
		if err != nil {
			return err
		}
	}
	log.Println("✅ All RabbitMQ exchanges setup completed")
	return nil
}

func SetupQueues(rmq *RabbitMQ, queueConfigs []QueueConfig) error {
	// Declare and bind all queues
	for _, q := range queueConfigs {
		_, err := rmq.DeclareQueue(q.Name, true, false)
		if err != nil {
			return err
		}

		err = rmq.BindQueue(q.Name, q.RoutingKey, q.Exchange)
		if err != nil {
			return err
		}
		log.Printf("✅ Queue '%s' bound to exchange '%s' with routing key '%s'", q.Name, q.Exchange, q.RoutingKey)
	}

	log.Println("✅ All RabbitMQ queues setup completed")
	return nil
}
