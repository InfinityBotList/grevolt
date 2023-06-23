package events

import "github.com/infinitybotlist/grevolt/types"

type MessageUpdate struct {
	Event

	// The message ID
	Id string `json:"id"`

	// Channel ID
	ChannelId string `json:"channel"`

	// Partial message object, not all data is available
	//
	// Exactly which fields are available is subject to change and not documented.
	Data *types.Message `json:"data"`
}
