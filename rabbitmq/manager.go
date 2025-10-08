package rabbitmq

import (
	"context"
	"log"
	"time"
)

// Manager quản lý publisher và consumer
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

// StartConsumers khởi động tất cả consumers cho profile service
func (m *Manager) StartConsumers(ctx context.Context) error {
	// Start consumer cho profile.created events
	err := m.consumer.ConsumeWithPrefetch(
		ctx,
		QueueProfileCreated,
		ConsumerProfileCreated,
		DefaultPrefetchCount,
		HandleProfileCreatedEvent,
	)
	if err != nil {
		return err
	}

	// Start consumer cho profile.updated events
	err = m.consumer.ConsumeWithPrefetch(
		ctx,
		QueueProfileUpdated,
		ConsumerProfileUpdated,
		DefaultPrefetchCount,
		HandleProfileUpdatedEvent,
	)
	if err != nil {
		return err
	}

	// Start consumer cho profile.deleted events
	err = m.consumer.ConsumeWithPrefetch(
		ctx,
		QueueProfileDeleted,
		ConsumerProfileDeleted,
		DefaultPrefetchCount,
		HandleProfileDeletedEvent,
	)
	if err != nil {
		return err
	}

	// Start consumer cho sharelink.created events
	err = m.consumer.ConsumeWithPrefetch(
		ctx,
		QueueShareLinkCreated,
		ConsumerShareLinkCreated,
		DefaultPrefetchCount,
		HandleShareLinkCreatedEvent,
	)
	if err != nil {
		return err
	}

	// Start consumer cho sharelink.accessed events
	err = m.consumer.ConsumeWithPrefetch(
		ctx,
		QueueShareLinkAccessed,
		ConsumerShareLinkAccessed,
		DefaultPrefetchCount,
		HandleShareLinkAccessedEvent,
	)
	if err != nil {
		return err
	}

	log.Println("✅ All RabbitMQ consumers started successfully")
	return nil
}

// PublishProfileCreated gửi event khi profile được tạo
func (m *Manager) PublishProfileCreated(ctx context.Context, userID, username, email string) error {
	event := ProfileCreatedEvent{
		UserID:    userID,
		Username:  username,
		Email:     email,
		CreatedAt: getCurrentTimestamp(),
	}

	return m.publisher.PublishMessage(ctx, ExchangeProfileEvents, RoutingKeyProfileCreated, event)
}

// PublishProfileUpdated gửi event khi profile được cập nhật
func (m *Manager) PublishProfileUpdated(ctx context.Context, userID, username string) error {
	event := ProfileUpdatedEvent{
		UserID:    userID,
		Username:  username,
		UpdatedAt: getCurrentTimestamp(),
	}

	return m.publisher.PublishMessage(ctx, ExchangeProfileEvents, RoutingKeyProfileUpdated, event)
}

// PublishProfileDeleted gửi event khi profile bị xóa
func (m *Manager) PublishProfileDeleted(ctx context.Context, userID string) error {
	event := ProfileDeletedEvent{
		UserID:    userID,
		DeletedAt: getCurrentTimestamp(),
	}

	return m.publisher.PublishMessage(ctx, ExchangeProfileEvents, RoutingKeyProfileDeleted, event)
}

// PublishShareLinkCreated gửi event khi share link được tạo
func (m *Manager) PublishShareLinkCreated(ctx context.Context, shareLinkID, userID, shortCode string) error {
	event := ShareLinkCreatedEvent{
		ShareLinkID: shareLinkID,
		UserID:      userID,
		ShortCode:   shortCode,
		CreatedAt:   getCurrentTimestamp(),
	}

	return m.publisher.PublishMessage(ctx, ExchangeShareLinkEvents, RoutingKeyShareLinkCreated, event)
}

// PublishShareLinkAccessed gửi event khi share link được truy cập
func (m *Manager) PublishShareLinkAccessed(ctx context.Context, shareLinkID, shortCode, ipAddress string) error {
	event := ShareLinkAccessedEvent{
		ShareLinkID: shareLinkID,
		ShortCode:   shortCode,
		AccessedAt:  getCurrentTimestamp(),
		IPAddress:   ipAddress,
	}

	return m.publisher.PublishMessage(ctx, ExchangeShareLinkEvents, RoutingKeyShareLinkAccessed, event)
}

// Helper function
func getCurrentTimestamp() int64 {
	return time.Now().Unix()
}
