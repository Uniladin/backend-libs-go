package rabbitmq

// ProfileCreatedEvent represents a profile creation event
type ProfileCreatedEvent struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"created_at"`
}

// ProfileUpdatedEvent represents a profile update event
type ProfileUpdatedEvent struct {
	UserID    string `json:"user_id"`
	Username  string `json:"username"`
	UpdatedAt int64  `json:"updated_at"`
}

// ProfileDeletedEvent represents a profile deletion event
type ProfileDeletedEvent struct {
	UserID    string `json:"user_id"`
	DeletedAt int64  `json:"deleted_at"`
}

// ShareLinkCreatedEvent represents a share link creation event
type ShareLinkCreatedEvent struct {
	ShareLinkID string `json:"share_link_id"`
	UserID      string `json:"user_id"`
	ShortCode   string `json:"short_code"`
	CreatedAt   int64  `json:"created_at"`
}

// ShareLinkAccessedEvent represents a share link access event
type ShareLinkAccessedEvent struct {
	ShareLinkID string `json:"share_link_id"`
	ShortCode   string `json:"short_code"`
	AccessedAt  int64  `json:"accessed_at"`
	IPAddress   string `json:"ip_address,omitempty"`
}
