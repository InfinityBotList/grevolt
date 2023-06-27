package types

// Representation of an Emoji on Revolt
type Emoji struct {
	// Unique Id
	Id string `json:"_id"`
	// What owns this emoji
	Parent *EmojiParent `json:"parent"`
	// Uploader user id
	CreatorId string `json:"creator_id"`
	// Emoji name
	Name string `json:"name"`
	// Whether the emoji is animated
	Animated bool `json:"animated,omitempty"`
	// Whether the emoji is marked as nsfw
	Nsfw bool `json:"nsfw,omitempty"`
}

type DataCreateEmoji struct {
	// Server name
	Name string `json:"name"`
	// Information about what owns this emoji
	Parent *EmojiParent `json:"parent"`
	// Whether the emoji is mature
	Nsfw bool `json:"nsfw,omitempty"`
}

type EmojiParent struct {
	// Type of emoji, either Server for server emoji or detached
	Type string `json:"type"`
	// ID of the server this emoji belongs to, if type is Server
	Id string `json:"id,omitempty"`
}
