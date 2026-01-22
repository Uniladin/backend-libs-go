package rabbitmq

// Exchange types
const (
	ExchangeTypeTopic  = "topic"  // Khớp theo mẫu (Wildcards) - Trung bình
	ExchangeTypeDirect = "direct" // Khớp chính xác từ khóa - Nhanh
	ExchangeTypeFanout = "fanout" // Gửi tất cả (Broadcast) - Rất nhanh
)

// Default settings
const (
	DefaultPrefetchCount = 10 // Nhận N messages cùng lúc
	DefaultMaxRetries    = 5  // Số lần retry tối đa
	DefaultRetryDelay    = 1  // seconds
)
