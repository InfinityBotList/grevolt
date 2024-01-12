package events

import "github.com/infinitybotlist/grevolt/types"

type ServerCreate struct {
	Event
	Server *types.Server `json:"server"`

	Channels []*types.Channel `json:"channels"`
	Emojis   []*types.Emoji   `json:"emojis"`
}
