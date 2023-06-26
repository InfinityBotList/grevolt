package events

import "github.com/infinitybotlist/grevolt/types"

// Work in progress, not well documented
type MessageAppendData struct {
	// Embeds in the message
	Embeds []*types.MessageEmbed `json:"embeds,omitempty"`
}

type MessageAppend struct {
	Event

	// The message ID
	Id string `json:"id"`

	// Channel ID
	ChannelId string `json:"channel"`

	// Partial message object, not all data is available
	//
	// Exactly which fields are available is subject to change and not documented.
	Append *MessageAppendData `json:"append"`
}
