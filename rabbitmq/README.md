# RabbitMQ Library

A reusable RabbitMQ client library for Go microservices.

## Features

- ✅ Connection management with auto-retry
- ✅ Publisher với confirmation support
- ✅ Consumer với ACK/NACK và QoS
- ✅ Predefined constants cho exchanges, queues, routing keys
- ✅ Event handlers
- ✅ High-level Manager API

## Installation

```bash
go get github.com/Uniladin/backend-libs-go/rabbitmq
```

## Usage

### Basic Connection

```go
import "github.com/Uniladin/backend-libs-go/rabbitmq"

// Configure
cfg := rabbitmq.Config{
    Addr: "localhost:5672",
    User: "guest",
    Pass: "guest",
}

// Connect
rmq, err := rabbitmq.NewRabbitMQ(cfg)
if err != nil {
    log.Fatal(err)
}
defer rmq.Close()
```

### Setup Exchanges and Queues

```go
err := rabbitmq.SetupProfileExchangeAndQueues(rmq)
if err != nil {
    log.Fatal(err)
}
```

### Publishing Events

```go
manager := rabbitmq.NewManager(rmq)

// Publish profile created event
err := manager.PublishProfileCreated(
    ctx,
    "user-123",
    "john_doe",
    "john@example.com",
)
```

### Consuming Events

```go
ctx := context.Background()

// Start all consumers
err := manager.StartConsumers(ctx)
if err != nil {
    log.Fatal(err)
}
```

## Constants

### Exchanges

- `ExchangeProfileEvents` - "profile.events"
- `ExchangeShareLinkEvents` - "sharelink.events"

### Queues

- `QueueProfileCreated` - "profile.created"
- `QueueProfileUpdated` - "profile.updated"
- `QueueProfileDeleted` - "profile.deleted"
- `QueueShareLinkCreated` - "sharelink.created"
- `QueueShareLinkAccessed` - "sharelink.accessed"

### Routing Keys

- `RoutingKeyProfileCreated` - "profile.created"
- `RoutingKeyProfileUpdated` - "profile.updated"
- `RoutingKeyProfileDeleted` - "profile.deleted"
- `RoutingKeyShareLinkCreated` - "sharelink.created"
- `RoutingKeyShareLinkAccessed` - "sharelink.accessed"

## Event Structures

```go
type EventProfileCreated struct {
    UserID    string `json:"user_id"`
    Username  string `json:"username"`
    Email     string `json:"email"`
    CreatedAt int64  `json:"created_at"`
}

type EventShareLinkCreated struct {
    ShareLinkID string `json:"share_link_id"`
    UserID      string `json:"user_id"`
    ShortCode   string `json:"short_code"`
    CreatedAt   int64  `json:"created_at"`
}
```

## Flow in service

- `setup.go`: Setup Exchange And Queues khi service start
- `handlers.go`: Xử lý các event nhận được từ RabbitMQ
- `manager.go`: Quản lý việc publish và consume các event

### Flow RabbitMQ

- publish message -> exchange -> routing -> queue -> consumer -> handler

## License

MIT
