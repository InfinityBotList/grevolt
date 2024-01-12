package events

type WebhookDelete struct {
	Event

	// Webhook ID
	Id string `json:"id"`
}
