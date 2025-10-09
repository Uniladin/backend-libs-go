package rabbitmq

// ======== AUTH EVENTS ========
type AuthCreatedEvent struct {
	UserID        string `json:"user_id"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}

// ======== END AUTH EVENTS ========

// ======== PROFILE EVENTS ========
type ProfileChangeEmailEvent struct {
	UserID   string `json:"user_id"`
	OldEmail string `json:"old_email"`
	NewEmail string `json:"new_email"`
}

type ProfileDeletedEvent struct {
	UserID string `json:"user_id"`
}

type ProfileUnDeletedEvent struct {
	UserID string `json:"user_id"`
}

// ======== END PROFILE EVENTS ========
