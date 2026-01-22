package profile

type ProfileChangeEmailEvent struct {
	UserID   string `json:"user_id"`
	OldEmail string `json:"old_email"`
	NewEmail string `json:"new_email"`
}

type ProfileDeletedEvent struct {
	UserID    string `json:"user_id"`
	DeletedAt int64  `json:"deleted_at"`
}

type ProfileUnDeletedEvent struct {
	UserID string `json:"user_id"`
}
