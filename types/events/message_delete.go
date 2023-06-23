package events

type MessageDelete struct {
	Event

	// The message ID
	Id string `json:"id"`

	// Channel ID
	ChannelId string `json:"channel"`

	// Undocumented but it exists?
	Ids []string `json:"ids,omitempty"`
}
