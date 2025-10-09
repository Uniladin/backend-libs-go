package rabbitmq

import "log"

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

// SetupProfileExchangeAndQueues sets up exchanges and queues for profile service
func SetupProfileExchangeAndQueues(rmq *RabbitMQ, exchangeConfigs []ExchangeConfig, queueConfigs []QueueConfig) error {
	for _, ec := range exchangeConfigs {
		err := rmq.DeclareExchange(ec.Name, ec.Type, true, false)
		if err != nil {
			return err
		}
	}

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

	log.Println("✅ All RabbitMQ exchanges and queues setup completed")
	return nil
}
