package events

import (
	"github.com/infinitybotlist/grevolt/types"
)

type Ready struct {
	Event
	Users    []*types.User    `json:"users,omitempty"`
	Servers  []*types.Server  `json:"servers,omitempty"`
	Channels []*types.Channel `json:"channels,omitempty"`
	Members  []*types.Member  `json:"members,omitempty"`
	Emojis   []*types.Emoji   `json:"emojis,omitempty"`
}
