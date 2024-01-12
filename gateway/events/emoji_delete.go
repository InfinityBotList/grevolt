package events

type EmojiDelete struct {
	Event

	// Emoji Id
	Id string `json:"id"`
}
