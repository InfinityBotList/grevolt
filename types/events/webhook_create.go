package events

import "github.com/infinitybotlist/grevolt/types"

type WebhookCreate struct {
	Event
	*types.Webhook
}
