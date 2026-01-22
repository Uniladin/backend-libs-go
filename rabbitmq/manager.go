package rabbitmq

// Manager quản lý publisher và consumer (GENERIC - dùng chung cho tất cả services)
type Manager struct {
	rabbitmq  *RabbitMQ
	publisher *Publisher
	consumer  *Consumer
}

// NewManager tạo RabbitMQ Manager mới
func NewManager(rabbitmq *RabbitMQ) *Manager {
	return &Manager{
		rabbitmq:  rabbitmq,
		publisher: NewPublisher(rabbitmq),
		consumer:  NewConsumer(rabbitmq),
	}
}

func (m *Manager) Close() {
	m.rabbitmq.Close()
}

// GetPublisher trả về publisher instance
func (m *Manager) GetPublisher() *Publisher {
	return m.publisher
}

// GetConsumer trả về consumer instance
func (m *Manager) GetConsumer() *Consumer {
	return m.consumer
}

// GetRabbitMQ trả về RabbitMQ connection instance
func (m *Manager) GetRabbitMQ() *RabbitMQ {
	return m.rabbitmq
}
