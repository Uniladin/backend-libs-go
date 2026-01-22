package auth

type EventAuthCreated struct {
	UserID        string `json:"user_id"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}
