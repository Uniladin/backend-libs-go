package rabbitmq

import "log"

// QueueConfig represents a queue configuration
type QueueConfig struct {
	Name       string
	RoutingKey string
	Exchange   string
}

// SetupProfileExchangeAndQueues sets up exchanges and queues for profile service
func SetupProfileExchangeAndQueues(rmq *RabbitMQ) error {
	// Declare main profile events exchange (topic type for flexible routing)
	err := rmq.DeclareExchange(ExchangeProfileEvents, ExchangeTypeTopic, true, false)
	if err != nil {
		return err
	}

	// Declare share link events exchange
	err = rmq.DeclareExchange(ExchangeShareLinkEvents, ExchangeTypeTopic, true, false)
	if err != nil {
		return err
	}

	// Declare queues for profile events
	profileQueues := []QueueConfig{
		{QueueProfileCreated, RoutingKeyProfileCreated, ExchangeProfileEvents},
		{QueueProfileUpdated, RoutingKeyProfileUpdated, ExchangeProfileEvents},
		{QueueProfileDeleted, RoutingKeyProfileDeleted, ExchangeProfileEvents},
	}

	for _, q := range profileQueues {
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

	// Declare queues for share link events
	shareLinkQueues := []QueueConfig{
		{QueueShareLinkCreated, RoutingKeyShareLinkCreated, ExchangeShareLinkEvents},
		{QueueShareLinkAccessed, RoutingKeyShareLinkAccessed, ExchangeShareLinkEvents},
	}

	for _, q := range shareLinkQueues {
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
