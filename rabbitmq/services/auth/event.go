package auth

type EventAuthCreated struct {
	UserID        string `json:"user_id"`
	Username      string `json:"username"`
	Phone         string `json:"phone,omitempty"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
}
