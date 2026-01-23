# RabbitMQ Library

Thư viện RabbitMQ dùng chung cho các microservice Go của Uniladin. Thư viện chuẩn hóa việc kết nối, publish/consume, và contract sự kiện giữa các service.

## Mục tiêu chính

- Kết nối RabbitMQ có auto-retry và tự khai báo exchange dùng chung.
- Publish/consume message theo JSON.
- ACK/NACK rõ ràng, hỗ trợ QoS (prefetch).
- Chuẩn hóa contract: exchange, routing key, queue, event.

## Cấu trúc thư mục

- `rabbitmq.go`: tạo kết nối, tự setup exchange theo danh sách ở [setup.go](setup.go).
- `setup.go`: danh sách exchange dùng chung và helper `SetupQueues`.
- `publisher.go`: publish message (có confirm).
- `consumer.go`: consume, ack/nack, QoS, helper `UnmarshalMessage`, `IsUnrecoverableError`.
- `services/<service>`: contract event/exchange/queue/routing của từng service.

## Cách tích hợp vào service mới (pkg/rabbitmq)

### 1) Thêm dependency

Nếu dùng submodule/local:

```go
replace uniladin-backend-libs => ./libs

require (
    uniladin-backend-libs v0.0.0-00010101000000-000000000000
)
```

Nếu dùng module:

```bash
go get uniladin-backend-libs
```

### 2) Khai báo contract trong libs (chia sẻ cho toàn hệ thống)

Tạo thư mục `libs/rabbitmq/services/<service>` và khai báo:

- `exchange.go`: tên exchange (vd: `foo.events`).
- `routing.go`: danh sách routing key theo event.
- `event.go`: struct event và JSON tags.

Ví dụ:

```go
package foo

const (
    ExchangeFooEvents     = "foo.events"
    RoutingKeyFooCreated  = "foo.created"
    RoutingKeyFooDeleted  = "foo.deleted"
)

type EventFooCreated struct {
    ID   string `json:"id"`
    Name string `json:"name"`
}
```

Sau đó cập nhật `libs/rabbitmq/setup.go` để đăng ký exchange mới:

```go
var exchangeConfigs = []ExchangeConfig{
    // ... existing exchanges
    {Name: foo.ExchangeFooEvents, Type: ExchangeTypeDirect},
}
```

### 3) Tạo `pkg/rabbitmq` trong service

Tạo tối thiểu các file:

- `constants.go`: tên queue và consumer tag riêng cho service.
- `setup.go`: khai báo queue config và gọi `rabbitmq.SetupQueues`.

Ví dụ `setup.go`:

```go
package rabbitmq

import (
    "uniladin-backend-libs/rabbitmq"
    "uniladin-backend-libs/rabbitmq/services/foo"
)

func SetupFooExchangeAndQueues(rmq *rabbitmq.RabbitMQ) error {
    queueConfigs := []rabbitmq.QueueConfig{
        {Name: QueueFooCreated, RoutingKey: foo.RoutingKeyFooCreated, Exchange: foo.ExchangeFooEvents},
        {Name: QueueFooDeleted, RoutingKey: foo.RoutingKeyFooDeleted, Exchange: foo.ExchangeFooEvents},
    }
    return rabbitmq.SetupQueues(rmq, queueConfigs)
}
```

### 4) Manager + handlers

- `manager.go`: wrap `rabbitmq.Manager`, khởi động consumer cần thiết.
- `handlers.go`: parse event bằng `rabbitmq.UnmarshalMessage`, gọi usecase.

Ví dụ start consumer:

```go
consumer := m.mqManager.GetConsumer()
return consumer.ConsumeWithPrefetch(
    ctx,
    QueueFooCreated,
    ConsumerFooCreated,
    rabbitmq.DefaultPrefetchCount,
    func(ctx context.Context, body []byte) error {
        var event foo.EventFooCreated
        if err := rabbitmq.UnmarshalMessage(body, &event); err != nil {
            return err
        }
        return m.deps.HandleFooCreated(ctx, event)
    },
)
```

### 5) Khởi tạo trong service

Khuyến nghị tạo hàm `Initialize` trong `pkg/rabbitmq/setup.go` để gom toàn bộ bước bootstrap:

```go
func Initialize(env *config.Env, deps Deps) *Manager {
    rmq, err := rabbitmq.NewRabbitMQ(rabbitmq.Config{
        Addr: env.RabbitMQAddr,
        User: env.RabbitMQUser,
        Pass: env.RabbitMQPass,
    })
    if err != nil {
        return nil
    }

    if err := SetupFooExchangeAndQueues(rmq); err != nil {
        return nil
    }

    libManager := rabbitmq.NewManager(rmq)
    serviceManager := NewManager(libManager, deps)
    if err := serviceManager.startConsumers(context.Background()); err != nil {
        return nil
    }

    return serviceManager
}
```

### 6) Publish event đến service khác

```go
publisher := rabbitmq.NewManager(rmq).GetPublisher()
err := publisher.PublishMessage(
    ctx,
    foo.ExchangeFooEvents,
    foo.RoutingKeyFooCreated,
    foo.EventFooCreated{ID: "123", Name: "demo"},
)
```

## Giải thích flow và ý nghĩa các khái niệm

### Event

- Là payload JSON mô tả một sự kiện nghiệp vụ (vd: `profile.deleted`).
- Được định nghĩa trong `libs/rabbitmq/services/<service>/event.go`.
- Là contract giữa các service: thay đổi cần cân nhắc tương thích.

### Exchange

- Điểm vào của message. Publisher gửi message vào exchange.
- Lib mặc định dùng `direct` (xem `setup.go`), nên routing key phải khớp chính xác.
- Có thể dùng `topic` cho wildcard hoặc `fanout` để broadcast (đã có const trong `constants.go`).

### Routing key

- Nhãn định tuyến cho từng event (vd: `profile.deleted`).
- Publisher gắn routing key, exchange dùng nó để route đến queue đã bind.

### Queue

- Nơi giữ message cho consumer.
- Mỗi service có thể có nhiều queue cho cùng exchange, tuỳ nhu cầu xử lý.
- Trong thực tế, tên queue thường nằm ở `pkg/rabbitmq/constants.go` để service tự quản.

### Consumer

- Tiến trình đọc queue và gọi handler.
- `ConsumeWithPrefetch` thiết lập QoS để giới hạn số message xử lý song song (`DefaultPrefetchCount`).
- ACK/NACK:
  - Handler thành công -> `Ack`.
  - Handler lỗi:
    - Nếu `IsUnrecoverableError` trả về true -> `Nack` và không requeue.
    - Ngược lại -> `Nack` và requeue để retry.

## Flow tổng thể

1. Service A tạo event (struct) và publish vào exchange kèm routing key.
2. Exchange định tuyến message đến queue đã bind bằng routing key.
3. Queue lưu message, consumer của service B đọc message.
4. Handler xử lý:
   - OK -> ACK, message bị xoá khỏi queue.
   - Lỗi -> NACK (requeue hoặc drop tuỳ loại lỗi).

