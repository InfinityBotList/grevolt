package events

type MessageReact struct {
	Event

	// The message ID
	Id string `json:"id"`

	// Channel ID
	ChannelId string `json:"channel_id"`

	// User ID <who made the reaction>
	UserId string `json:"user_id"`

	// Emoji ID <of the reaction>
	EmojiId string `json:"emoji_id"`
}
