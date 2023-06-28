package types

// WebhookList : List of webhooks
type WebhookList []*Webhook

// FieldsWebhook : Optional fields on webhook object
//
// <undocumented, from https://github.com/revoltchat/backend/blob/master/crates/core/database/src/models/channel_webhooks/ops/mongodb.rs#L71>
type FieldsWebhook string

const (
	AVATAR_FieldsWebhook FieldsWebhook = "Avatar"
)

// Webhook
type Webhook struct {
	// Webhook Id
	Id string `json:"id"`
	// The name of the webhook
	Name string `json:"name"`
	// The avatar of the webhook
	Avatar *File `json:"avatar,omitempty"`
	// The channel this webhook belongs to
	ChannelId string `json:"channel_id"`
	// The private token for the webhook
	Token string `json:"token,omitempty"`
}

// Data for creating a webhook
type DataCreateWebhook struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar,omitempty"`
}
