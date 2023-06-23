package events

import (
	"github.com/infinitybotlist/grevolt/types"
)

type Ready struct {
	Event
	Users    []*types.User    `json:"users"`
	Servers  []*types.Server  `json:"servers"`
	Channels []*types.Channel `json:"channels"`
	Members  []*types.Member  `json:"members"`
	Emojis   []*types.Emoji   `json:"emojis"`
}
