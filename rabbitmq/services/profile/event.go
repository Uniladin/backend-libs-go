package profile

type EventProfileChangeEmail struct {
	UserID   string `json:"user_id"`
	OldEmail string `json:"old_email"`
	NewEmail string `json:"new_email"`
}

type EventProfileDeleted struct {
	UserID    string `json:"user_id"`
	DeletedAt int64  `json:"deleted_at"`
}

type EventProfileUnDeleted struct {
	UserID string `json:"user_id"`
}
