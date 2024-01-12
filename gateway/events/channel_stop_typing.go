package events

type ChannelStopTyping struct {
	Event

	// Channel Id
	Id string `json:"id"`

	// User Id
	UserId string `json:"user"`
}
