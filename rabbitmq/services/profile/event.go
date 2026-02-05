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

type EventProfileIncreaseElo struct {
	PersonaID     string `json:"persona_id"`
	ActionType    string `json:"action_type"`
	Source        string `json:"source"`
	DeltaOverride *int32 `json:"delta_override,omitempty"`
}
