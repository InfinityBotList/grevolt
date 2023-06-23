package events

type MessageRemoveReaction struct {
	Event

	// The message ID
	Id string `json:"id"`

	// Channel ID
	ChannelId string `json:"channel_id"`

	// Emoji ID <of the reaction>
	EmojiId string `json:"emoji_id"`
}
