package types

// Fetch settings from server filtered by keys.
type DataFetchSettings struct {
	// Keys to fetch
	Keys []string `json:"keys"`
}

// Composite key pointing to a user's view of a channel
type ChannelUnreadId struct {
	// Channel Id
	Channel string `json:"channel"`
	// User Id
	User string `json:"user"`
}

type UnreadMessage struct {
	// Composite primary key consisting of channel and user id
	Id *ChannelUnreadId `json:"_id"`
	// Id of the last message read in this channel by a user
	LastId string `json:"last_id,omitempty"`
	// Array of message ids that mention the user
	Mentions []string `json:"mentions,omitempty"`
}
