package rabbitmq

// =============== EXCHANGE NAMES ==============

// Exchange names
const (
	ExchangeAuthEvents    = "auth.events"
	ExchangeProfileEvents = "profile.events"
)

// Exchange types
const (
	ExchangeTypeTopic  = "topic"
	ExchangeTypeDirect = "direct"
	ExchangeTypeFanout = "fanout"
)

// =============== END EXCHANGE NAMES ==============

// =============== QUEUE NAMES ==============

// Queue names - Auth
const (
	QueueAuthCreated = "auth.created"
)

// Queue names - Profile
const (
	QueueProfileChangeEmail = "profile.change_email"
	QueueProfileDeleted     = "profile.deleted"
)

// =============== END QUEUE NAMES ==============

// =============== ROUTING KEYS ==============

// Routing keys - Auth
const (
	RoutingKeyAuthCreated = "auth.created"
)

// Routing keys - Profile
const (
	RoutingKeyProfileChangeEmail = "profile.change_email"
	RoutingKeyProfileDeleted     = "profile.deleted"
	RoutingKeyProfileUndeleted   = "profile.undeleted"
)

// =============== END ROUTING KEYS ==============

// Default settings
const (
	DefaultPrefetchCount = 10 // Nhận N messages cùng lúc
	DefaultMaxRetries    = 5  // Số lần retry tối đa
	DefaultRetryDelay    = 1  // seconds
)
