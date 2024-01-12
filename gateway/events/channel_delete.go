package events

type ChannelDelete struct {
	Event

	// Channel Id
	Id string `json:"id"`
}
