package rabbitmq

// =============== EXCHANGE NAMES ==============

// Exchange names
const (
	ExchangeAuthEvents         = "auth.events"
	ExchangeProfileEvents      = "profile.events"
	ExchangeNotificationEvents = "notification.events"
)

// Exchange types
const (
	ExchangeTypeTopic  = "topic"  // Khớp theo mẫu (Wildcards) - Trung bình
	ExchangeTypeDirect = "direct" // Khớp chính xác từ khóa - Nhanh
	ExchangeTypeFanout = "fanout" // Gửi tất cả (Broadcast) - Rất nhanh
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

// Queue names - Notification
const (
	QueueNotificationLiked       = "notification.liked"
	QueueNotificationNewComment  = "notification.new_comment"
	QueueNotificationViewProfile = "notification.view_profile"
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

// Routing keys - Notification
const (
	RoutingKeyNotificationLiked       = "notification.liked"
	RoutingKeyNotificationNewComment  = "notification.new_comment"
	RoutingKeyNotificationViewProfile = "notification.view_profile"
)

// =============== END ROUTING KEYS ==============

// Default settings
const (
	DefaultPrefetchCount = 10 // Nhận N messages cùng lúc
	DefaultMaxRetries    = 5  // Số lần retry tối đa
	DefaultRetryDelay    = 1  // seconds
)
