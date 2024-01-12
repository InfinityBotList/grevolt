package events

import "github.com/infinitybotlist/grevolt/types"

type EmojiCreate struct {
	Event
	*types.Emoji
}
