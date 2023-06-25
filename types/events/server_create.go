package events

import "github.com/infinitybotlist/grevolt/types"

type ServerCreate struct {
	Event
	*types.Server

	// <undocumented, may exist???>
	Channels []*types.Channel `json:"channels,omitempty"`
}
