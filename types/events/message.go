package events

import "github.com/infinitybotlist/grevolt/types"

type Message struct {
	Event
	*types.Message
}
