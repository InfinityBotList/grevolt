package events

import "github.com/infinitybotlist/grevolt/types"

type ChannelCreate struct {
	Event
	*types.Channel
}
