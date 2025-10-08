package rabbitmq

// =============== EXCHANGE NAMES ==============

// Exchange names
const (
	ExchangeProfileEvents   = "profile.events"
	ExchangeShareLinkEvents = "sharelink.events"
)

// Exchange types
const (
	ExchangeTypeTopic  = "topic"
	ExchangeTypeDirect = "direct"
	ExchangeTypeFanout = "fanout"
)

// =============== END EXCHANGE NAMES ==============

// =============== QUEUE NAMES ==============

// Queue names - Profile
const (
	QueueProfileCreated = "profile.created"
	QueueProfileUpdated = "profile.updated"
	QueueProfileDeleted = "profile.deleted"
)

// Queue names - ShareLink
const (
	QueueShareLinkCreated  = "sharelink.created"
	QueueShareLinkAccessed = "sharelink.accessed"
)

// =============== END QUEUE NAMES ==============

// =============== ROUTING KEYS ==============

// Routing keys - Profile
const (
	RoutingKeyProfileCreated = "profile.created"
	RoutingKeyProfileUpdated = "profile.updated"
	RoutingKeyProfileDeleted = "profile.deleted"
)

// Routing keys - ShareLink
const (
	RoutingKeyShareLinkCreated  = "sharelink.created"
	RoutingKeyShareLinkAccessed = "sharelink.accessed"
)

// =============== END ROUTING KEYS ==============

// Consumer tags
const (
	ConsumerProfileCreated    = "profile-service-created-consumer"
	ConsumerProfileUpdated    = "profile-service-updated-consumer"
	ConsumerProfileDeleted    = "profile-service-deleted-consumer"
	ConsumerShareLinkCreated  = "sharelink-service-created-consumer"
	ConsumerShareLinkAccessed = "sharelink-service-accessed-consumer"
)

// Default settings
const (
	DefaultPrefetchCount = 10
	DefaultMaxRetries    = 5
	DefaultRetryDelay    = 1 // seconds
)
