# Uniladin Backend Libraries (Go)

Shared libraries for Uniladin microservices ecosystem.

## 📦 Packages

### RabbitMQ (`/rabbitmq`)

Complete RabbitMQ client library with connection management, publishing, consuming, and predefined constants.

**Features:**

- ✅ Connection management với auto-retry
- ✅ Publisher với confirmation support
- ✅ Consumer với ACK/NACK và QoS
- ✅ Predefined constants (exchanges, queues, routing keys)
- ✅ Event handlers
- ✅ High-level Manager API

**Installation:**

```go
import "github.com/Uniladin/backend-libs-go/rabbitmq"
```

**Quick Start:**

```go
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

// Setup
rabbitmq.SetupProfileExchangeAndQueues(rmq)

// Use manager
manager := rabbitmq.NewManager(rmq)
manager.PublishProfileCreated(ctx, userID, username, email)
```

**Documentation:** See [rabbitmq/README.md](rabbitmq/README.md)

## 🔧 Usage in Projects

### As Git Submodule (Recommended for Development)

```bash
# Add submodule
git submodule add https://github.com/Uniladin/backend-libs-go.git libs

# Update go.mod
```

**go.mod:**

```go
module your-service

replace github.com/Uniladin/backend-libs-go => ./libs

require (
    github.com/Uniladin/backend-libs-go v0.0.0-00010101000000-000000000000
)
```

### As Go Module (Production)

```bash
go get github.com/Uniladin/backend-libs-go
```

**go.mod:**

```go
require (
    github.com/Uniladin/backend-libs-go v1.0.0
)
```

## 📚 Constants Reference

All packages provide typed constants to avoid hardcoded strings:

```go
// RabbitMQ exchanges
rabbitmq.ExchangeProfileEvents   // "profile.events"
rabbitmq.ExchangeShareLinkEvents // "sharelink.events"

// Queues
rabbitmq.QueueProfileCreated // "profile.created"
rabbitmq.QueueProfileUpdated // "profile.updated"

// Routing keys
rabbitmq.RoutingKeyProfileCreated // "profile.created"

// Default settings
rabbitmq.DefaultPrefetchCount // 10
```

## 🚀 Services Using This Library

- ✅ backend-profile
- ⏳ backend-chat (coming soon)
- ⏳ backend-notification (coming soon)
- ⏳ backend-payment (coming soon)

## 🔄 Development Workflow

### Making Changes

1. **Edit libs code:**

   ```bash
   cd libs
   # Make changes to rabbitmq/ or other packages
   ```

2. **Test in service:**

   ```bash
   cd ../
   go test ./...
   go build ./...
   ```

3. **Commit libs changes:**

   ```bash
   cd libs
   git add .
   git commit -m "Add feature X"
   git push origin main
   ```

4. **Update service to use new libs:**
   ```bash
   cd ../
   git add libs
   git commit -m "Update libs submodule"
   ```

### Updating Libs in Other Services

```bash
# In any service using libs
git submodule update --remote libs
git add libs
git commit -m "Update libs to latest"
```

## 📖 Documentation

Each package has its own README:

- [rabbitmq/README.md](rabbitmq/README.md) - RabbitMQ client library

## 🤝 Contributing

1. Create feature branch
2. Make changes
3. Test in at least one service
4. Create PR
5. After merge, update services

## 📝 License

MIT

## 🔗 Links

- GitHub: https://github.com/Uniladin/backend-libs-go
- Issues: https://github.com/Uniladin/backend-libs-go/issues
