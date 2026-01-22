package notification

type EventSendNotificationToUser struct {
	AccountID string            `json:"account_id"`
	Title     string            `json:"title,omitempty"`
	Body      string            `json:"body,omitempty"`
	ImageURL  string            `json:"image_url,omitempty"`
	Data      map[string]string `json:"data,omitempty"`
}

type EventSendNotificationToToken struct {
	FcmToken string            `json:"fcm_token"`
	Title    string            `json:"title,omitempty"`
	Body     string            `json:"body,omitempty"`
	ImageURL string            `json:"image_url,omitempty"`
	Data     map[string]string `json:"data,omitempty"`
}

type EventSendNotificationToTopic struct {
	Topic    string            `json:"topic"`
	Title    string            `json:"title,omitempty"`
	Body     string            `json:"body,omitempty"`
	ImageURL string            `json:"image_url,omitempty"`
	Data     map[string]string `json:"data,omitempty"`
}
